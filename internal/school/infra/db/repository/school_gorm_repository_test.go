package school_repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
	school_model "github.com/williamkoller/system-education/internal/school/infra/db/model"
	port_school_repository "github.com/williamkoller/system-education/internal/school/port/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SchoolGormRepositorySuite struct {
	suite.Suite
	db         *gorm.DB
	repository *SchoolGormRepository
}

func (s *SchoolGormRepositorySuite) SetupTest() {
	s.db = setupTestDB(s.T())
	s.repository = NewSchoolGormRepository(s.db)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&school_model.School{})
	assert.NoError(t, err)

	return db
}

func (s *SchoolGormRepositorySuite) TestCreate() {
	school := &school_entity.School{
		ID:          "school-1",
		Name:        "Test School",
		Code:        "TS001",
		Address:     "123 Test St",
		City:        "Test City",
		State:       "TS",
		ZipCode:     "12345",
		Country:     "Test Country",
		PhoneNumber: "123-456-7890",
		Email:       "test@school.com",
		IsActive:    true,
		Description: "A test school",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdSchool, err := s.repository.Save(school)

	s.NoError(err)
	s.NotNil(createdSchool)
	s.Equal(school.ID, createdSchool.ID)
	s.Equal(school.Name, createdSchool.Name)
	s.Equal(school.Code, createdSchool.Code)
}

func (s *SchoolGormRepositorySuite) TestCreate_Error() {
	school := &school_entity.School{
		ID:   "school-2",
		Name: "Test School 2",
	}

	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	createdSchool, err := s.repository.Save(school)

	s.Error(err)
	s.Nil(createdSchool)
}

func (s *SchoolGormRepositorySuite) TestUpdate() {
	school := &school_entity.School{
		ID:        "school-3",
		Name:      "Update School",
		Code:      "US001",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	created, _ := s.repository.Save(school)

	created.Name = "Updated School Name"
	updatedSchool, err := s.repository.Update(created.ID, created)

	s.NoError(err)
	s.NotNil(updatedSchool)
	s.Equal("Updated School Name", updatedSchool.Name)

	found, _ := s.repository.FindById(created.ID)
	s.Equal("Updated School Name", found.Name)
}

func (s *SchoolGormRepositorySuite) TestUpdate_NotFound() {
	school := &school_entity.School{
		ID:   "school-4",
		Name: "Not Found School",
	}

	updatedSchool, err := s.repository.Update("non-existent-id", school)

	s.Error(err)
	s.Equal(port_school_repository.ErrNotFound, err)
	s.Nil(updatedSchool)
}

func (s *SchoolGormRepositorySuite) TestUpdate_DBError() {
	school := &school_entity.School{
		ID:        "school-5",
		Name:      "DB Error School",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	created, _ := s.repository.Save(school)

	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	updatedSchool, err := s.repository.Update(created.ID, created)

	s.Error(err)
	s.NotEqual(port_school_repository.ErrNotFound, err)
	s.Nil(updatedSchool)
}

func (s *SchoolGormRepositorySuite) TestDelete() {
	school := &school_entity.School{
		ID:        "school-6",
		Name:      "Delete School",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	created, _ := s.repository.Save(school)

	err := s.repository.Delete(created.ID)

	s.NoError(err)

	found, err := s.repository.FindById(created.ID)
	s.Error(err)
	s.Equal(port_school_repository.ErrNotFound, err)
	s.Nil(found)
}

func (s *SchoolGormRepositorySuite) TestDelete_NotFound() {
	err := s.repository.Delete("non-existent-id")

	s.Error(err)
	s.Equal(port_school_repository.ErrNotFound, err)
}

func (s *SchoolGormRepositorySuite) TestDelete_DBError() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	err := s.repository.Delete("some-id")

	s.Error(err)
	s.NotEqual(port_school_repository.ErrNotFound, err)
}

func (s *SchoolGormRepositorySuite) TestFindAll() {
	school1 := &school_entity.School{ID: "school-7", Name: "School 1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	school2 := &school_entity.School{ID: "school-8", Name: "School 2", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	s.repository.Save(school1)
	s.repository.Save(school2)

	schools, err := s.repository.FindAll()

	s.NoError(err)
	s.Len(schools, 2)
}

func (s *SchoolGormRepositorySuite) TestFindAll_Error() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	schools, err := s.repository.FindAll()

	s.Error(err)
	s.Nil(schools)
}

func (s *SchoolGormRepositorySuite) TestFindById() {
	school := &school_entity.School{
		ID:        "school-9",
		Name:      "Find By ID School",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	created, _ := s.repository.Save(school)

	foundSchool, err := s.repository.FindById(created.ID)

	s.NoError(err)
	s.NotNil(foundSchool)
	s.Equal(created.ID, foundSchool.ID)
	s.Equal(created.Name, foundSchool.Name)
}

func (s *SchoolGormRepositorySuite) TestFindById_NotFound() {
	foundSchool, err := s.repository.FindById("non-existent-id")

	s.Error(err)
	s.Equal(port_school_repository.ErrNotFound, err)
	s.Nil(foundSchool)
}

func (s *SchoolGormRepositorySuite) TestFindById_DBError() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	foundSchool, err := s.repository.FindById("some-id")

	s.Error(err)
	s.NotEqual(port_school_repository.ErrNotFound, err)
	s.Nil(foundSchool)
}

func TestSchoolGormRepositorySuite(t *testing.T) {
	suite.Run(t, new(SchoolGormRepositorySuite))
}
