package student_usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	student_usecase "github.com/williamkoller/system-education/internal/student/application/usecase"
	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
	port_student_repository "github.com/williamkoller/system-education/internal/student/port/repository"
	student_dtos "github.com/williamkoller/system-education/internal/student/presentation/dtos"
)

type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) Save(ctx context.Context, s *student_entity.Student) (*student_entity.Student, error) {
	args := m.Called(ctx, s)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*student_entity.Student), args.Error(1)
}

func (m *MockStudentRepository) FindAll(ctx context.Context) ([]*student_entity.Student, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*student_entity.Student), args.Error(1)
}

func (m *MockStudentRepository) FindById(ctx context.Context, id string) (*student_entity.Student, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*student_entity.Student), args.Error(1)
}

func (m *MockStudentRepository) Update(ctx context.Context, id string, s *student_entity.Student) (*student_entity.Student, error) {
	args := m.Called(ctx, id, s)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*student_entity.Student), args.Error(1)
}

func (m *MockStudentRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestStudentUsecase_Create(t *testing.T) {
	t.Run("should create student successfully", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()

		input := student_dtos.AddStudentDto{
			FullName:       "John Doe",
			EnrollmentCode: "2023001",
			Email:          "john@example.com",
			PhoneNumber:    "1234567890",
			DateOfBirth:    time.Now().AddDate(-10, 0, 0),
			CPF:            "97093236014", // Valid CPF
			RG:             "1234567",
			Address:        "123 Main St",
			City:           "City",
			State:          "ST",
			ZipCode:        "12345",
			Country:        "Country",
			SchoolID:       "school-1",
			SchoolName:     "School Name",
			SchoolCode:     "SC001",
			Grade:          "5th",
			ClassRoom:      "A",
			Shift:          "morning",
			EnrollmentDate: time.Now(),
			GuardianName:   "Guardian",
			GuardianPhone:  "0987654321",
			GuardianEmail:  "guardian@example.com",
			GuardianCPF:    "97093236014", // Valid CPF
			IsActive:       true,
			Observations:   "None",
		}

		mockRepo.On("Save", ctx, mock.AnythingOfType("*student_entity.Student")).Return(&student_entity.Student{
			ID: "generated-id",
			PersonalInfo: student_entity.PersonalInfo{
				FullName: input.FullName,
				Email:    input.Email,
			},
		}, nil)

		result, err := usecase.Create(ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "generated-id", result.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when validation fails", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()

		input := student_dtos.AddStudentDto{
			FullName: "", // Invalid: empty name
		}

		result, err := usecase.Create(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertNotCalled(t, "Save")
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()

		input := student_dtos.AddStudentDto{
			FullName:       "John Doe",
			EnrollmentCode: "2023001",
			Email:          "john@example.com",
			PhoneNumber:    "1234567890",
			DateOfBirth:    time.Now().AddDate(-10, 0, 0),
			CPF:            "97093236014",
			RG:             "1234567",
			Address:        "123 Main St",
			City:           "City",
			State:          "ST",
			ZipCode:        "12345",
			Country:        "Country",
			SchoolID:       "school-1",
			SchoolName:     "School Name",
			SchoolCode:     "SC001",
			Grade:          "5th",
			ClassRoom:      "A",
			Shift:          "morning",
			EnrollmentDate: time.Now(),
			GuardianName:   "Guardian",
			GuardianPhone:  "0987654321",
			GuardianEmail:  "guardian@example.com",
			GuardianCPF:    "97093236014",
			IsActive:       true,
			Observations:   "None",
		}

		mockRepo.On("Save", ctx, mock.Anything).Return((*student_entity.Student)(nil), errors.New("db error"))

		result, err := usecase.Create(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentUsecase_FindAll(t *testing.T) {
	t.Run("should return all students", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()

		expectedStudents := []*student_entity.Student{
			{ID: "1", PersonalInfo: student_entity.PersonalInfo{FullName: "Student 1"}},
			{ID: "2", PersonalInfo: student_entity.PersonalInfo{FullName: "Student 2"}},
		}

		mockRepo.On("FindAll", ctx).Return(expectedStudents, nil)

		result, err := usecase.FindAll(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, expectedStudents, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()

		mockRepo.On("FindAll", ctx).Return(([]*student_entity.Student)(nil), errors.New("db error"))

		result, err := usecase.FindAll(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentUsecase_FindById(t *testing.T) {
	t.Run("should return student by id", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()
		id := "123"

		expectedStudent := &student_entity.Student{ID: id, PersonalInfo: student_entity.PersonalInfo{FullName: "Student 1"}}

		mockRepo.On("FindById", ctx, id).Return(expectedStudent, nil)

		result, err := usecase.FindById(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, expectedStudent, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when student not found", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()
		id := "123"

		mockRepo.On("FindById", ctx, id).Return((*student_entity.Student)(nil), port_student_repository.ErrNotFound)

		result, err := usecase.FindById(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, port_student_repository.ErrNotFound, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestStudentUsecase_Update(t *testing.T) {
	// Setup
	mockRepo := new(MockStudentRepository)
	usecase := student_usecase.NewStudentUsecase(mockRepo)
	ctx := context.Background()

	// Data
	studentID := "student-123"
	validCPF := "97093236014"
	existingStudent := &student_entity.Student{
		ID: studentID,
		PersonalInfo: student_entity.PersonalInfo{
			FullName:    "John Doe",
			Email:       "john@example.com",
			CPF:         validCPF,
			DateOfBirth: time.Now().AddDate(-10, 0, 0),
		},
		Address: student_entity.AddressInfo{
			Address: "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
			Country: "USA",
		},
		School: student_entity.SchoolInfo{
			SchoolID: "school-1",
			Shift:    student_entity.StudentShiftMorning,
		},
		Guardian: student_entity.GuardianInfo{
			Name: "Jane Doe",
			CPF:  validCPF,
		},
		IsActive: true,
	}

	newFullName := "John Updated"
	newEmail := "john.updated@example.com"
	updateDto := student_dtos.UpdateStudentDto{
		FullName: &newFullName,
		Email:    &newEmail,
	}

	t.Run("should update student successfully", func(t *testing.T) {
		mockRepo.On("FindById", ctx, studentID).Return(existingStudent, nil)
		mockRepo.On("Update", ctx, studentID, mock.Anything).Return(existingStudent, nil)

		updatedStudent, err := usecase.Update(ctx, studentID, updateDto)

		assert.NoError(t, err)
		assert.NotNil(t, updatedStudent)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when student not found", func(t *testing.T) {
		mockRepo.On("FindById", ctx, "unknown").Return((*student_entity.Student)(nil), port_student_repository.ErrNotFound)

		result, err := usecase.Update(ctx, "unknown", updateDto)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestStudentUsecase_Delete(t *testing.T) {
	t.Run("should delete student successfully", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()
		id := "123"

		mockRepo.On("Delete", ctx, id).Return(nil)

		err := usecase.Delete(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when delete fails", func(t *testing.T) {
		mockRepo := new(MockStudentRepository)
		usecase := student_usecase.NewStudentUsecase(mockRepo)
		ctx := context.Background()
		id := "123"

		mockRepo.On("Delete", ctx, id).Return(errors.New("delete error"))

		err := usecase.Delete(ctx, id)

		assert.Error(t, err)
		assert.Equal(t, "delete error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
