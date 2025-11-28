package student_repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	school_model "github.com/williamkoller/system-education/internal/school/infra/db/model"
	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
	student_model "github.com/williamkoller/system-education/internal/student/infra/db/model"
	port_student_repository "github.com/williamkoller/system-education/internal/student/port/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentGormRepositorySuite struct {
	suite.Suite
	db         *gorm.DB
	repository *StudentGormRepository
}

func (s *StudentGormRepositorySuite) SetupTest() {
	s.db = setupTestDB(s.T())
	s.repository = NewStudentGormRepository(s.db)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&school_model.School{}, &student_model.Student{})
	assert.NoError(t, err)

	return db
}

func TestStudentGormRepositorySuite(t *testing.T) {
	suite.Run(t, new(StudentGormRepositorySuite))
}

func createValidStudent() *student_entity.Student {
	return &student_entity.Student{
		ID: "student-1",
		PersonalInfo: student_entity.PersonalInfo{
			FullName:       "John Doe",
			EnrollmentCode: "ST123",
			Email:          "john@example.com",
			PhoneNumber:    "123456789",
			DateOfBirth:    time.Now().AddDate(-10, 0, 0),
			CPF:            "111.444.777-35",
			RG:             "RG123",
		},
		Address: student_entity.AddressInfo{
			Address: "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
			Country: "USA",
		},
		School: student_entity.SchoolInfo{
			SchoolID:       "school-1",
			SchoolName:     "Test School",
			SchoolCode:     "TS",
			Grade:          "5",
			ClassRoom:      "A",
			Shift:          student_entity.StudentShiftMorning,
			EnrollmentDate: time.Now(),
		},
		Guardian: student_entity.GuardianInfo{
			Name:  "Jane Doe",
			Phone: "987654321",
			Email: "jane@example.com",
			CPF:   "111.444.777-35",
		},
		IsActive:     true,
		Observations: "Test observation",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (s *StudentGormRepositorySuite) TestSave() {
	// Create a school first to satisfy foreign key if needed (though SQLite memory might be loose, better safe)
	school := &school_model.School{
		ID:   "school-1",
		Name: "Test School",
		Code: "TS",
	}
	s.db.Create(school)

	student := createValidStudent()

	savedStudent, err := s.repository.Save(context.Background(), student)

	s.NoError(err)
	s.NotNil(savedStudent)
	s.Equal(student.ID, savedStudent.ID)
	s.Equal(student.PersonalInfo.FullName, savedStudent.PersonalInfo.FullName)
	s.Equal(student.PersonalInfo.EnrollmentCode, savedStudent.PersonalInfo.EnrollmentCode)
}

func (s *StudentGormRepositorySuite) TestSave_Error() {
	student := createValidStudent()

	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	savedStudent, err := s.repository.Save(context.Background(), student)

	s.Error(err)
	s.Nil(savedStudent)
}

func (s *StudentGormRepositorySuite) TestFindAll() {
	student1 := createValidStudent()
	student1.ID = "student-1"
	student1.PersonalInfo.EnrollmentCode = "ST1"
	student1.PersonalInfo.Email = "student1@example.com"
	student1.PersonalInfo.CPF = "111.111.111-11"

	student2 := createValidStudent()
	student2.ID = "student-2"
	student2.PersonalInfo.EnrollmentCode = "ST2"
	student2.PersonalInfo.Email = "student2@example.com"
	student2.PersonalInfo.CPF = "222.222.222-22"

	s.repository.Save(context.Background(), student1)
	s.repository.Save(context.Background(), student2)

	students, err := s.repository.FindAll(context.Background())

	s.NoError(err)
	s.Len(students, 2)
}

func (s *StudentGormRepositorySuite) TestFindById() {
	student := createValidStudent()
	created, _ := s.repository.Save(context.Background(), student)

	foundStudent, err := s.repository.FindById(context.Background(), created.ID)

	s.NoError(err)
	s.NotNil(foundStudent)
	s.Equal(created.ID, foundStudent.ID)
	s.Equal(created.PersonalInfo.FullName, foundStudent.PersonalInfo.FullName)
}

func (s *StudentGormRepositorySuite) TestFindById_NotFound() {
	foundStudent, err := s.repository.FindById(context.Background(), "non-existent-id")

	s.Error(err)
	s.Equal(port_student_repository.ErrNotFound, err)
	s.Nil(foundStudent)
}

func (s *StudentGormRepositorySuite) TestUpdate() {
	student := createValidStudent()
	created, _ := s.repository.Save(context.Background(), student)

	created.PersonalInfo.FullName = "Updated Name"
	updatedStudent, err := s.repository.Update(context.Background(), created.ID, created)

	s.NoError(err)
	s.NotNil(updatedStudent)
	s.Equal("Updated Name", updatedStudent.PersonalInfo.FullName)

	found, _ := s.repository.FindById(context.Background(), created.ID)
	s.Equal("Updated Name", found.PersonalInfo.FullName)
}

func (s *StudentGormRepositorySuite) TestUpdate_NotFound() {
	student := createValidStudent()
	updatedStudent, err := s.repository.Update(context.Background(), "non-existent-id", student)

	s.Error(err)
	s.Equal(port_student_repository.ErrNotFound, err)
	s.Nil(updatedStudent)
}

func (s *StudentGormRepositorySuite) TestDelete() {
	student := createValidStudent()
	created, _ := s.repository.Save(context.Background(), student)

	err := s.repository.Delete(context.Background(), created.ID)

	s.NoError(err)

	found, err := s.repository.FindById(context.Background(), created.ID)
	s.Error(err)
	s.Equal(port_student_repository.ErrNotFound, err)
	s.Nil(found)
}

func (s *StudentGormRepositorySuite) TestDelete_NotFound() {
	err := s.repository.Delete(context.Background(), "non-existent-id")

	s.Error(err)
	s.Equal(port_student_repository.ErrNotFound, err)
}
