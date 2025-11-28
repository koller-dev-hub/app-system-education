package student_usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	student_usecase "github.com/williamkoller/system-education/internal/student/application/usecase"
	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
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

func TestStudentUsecase_Update(t *testing.T) {
	// Setup
	mockRepo := new(MockStudentRepository)
	usecase := student_usecase.NewStudentUsecase(mockRepo)
	ctx := context.Background()

	// Data
	studentID := "student-123"
	validCPF := "97093236014" // Valid CPF for testing
	existingStudent := &student_entity.Student{
		ID: studentID,
		PersonalInfo: student_entity.PersonalInfo{
			FullName:    "John Doe",
			Email:       "john@example.com",
			CPF:         validCPF,
			DateOfBirth: time.Now().AddDate(-10, 0, 0), // 10 years old
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

	// Expectations
	mockRepo.On("FindById", ctx, studentID).Return(existingStudent, nil)

	// We expect Update to be called with the student having updated fields
	// Note: Validation might format CPF, so we should expect formatted or handle it.
	// The existingStudent already has unformatted CPF in this setup.
	// ValidationUpdateStudent calls utils.FormatCPF if valid.
	// So in Update expectation, CPF might be formatted if validation passes.
	// Let's check if ValidationUpdateStudent formats it in place. Yes it does: s.PersonalInfo.CPF = utils.FormatCPF(s.PersonalInfo.CPF)

	mockRepo.On("Update", ctx, studentID, mock.MatchedBy(func(s *student_entity.Student) bool {
		return s.PersonalInfo.FullName == newFullName && s.PersonalInfo.Email == newEmail
	})).Return(existingStudent, nil)

	// Execute
	updatedStudent, err := usecase.Update(ctx, studentID, updateDto)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, updatedStudent)
	assert.Equal(t, newFullName, updatedStudent.PersonalInfo.FullName)
	assert.Equal(t, newEmail, updatedStudent.PersonalInfo.Email)

	mockRepo.AssertExpectations(t)
}

func TestStudentUsecase_Update_Partial(t *testing.T) {
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
	// Email is NOT updated
	updateDto := student_dtos.UpdateStudentDto{
		FullName: &newFullName,
	}

	// Expectations
	mockRepo.On("FindById", ctx, studentID).Return(existingStudent, nil)

	mockRepo.On("Update", ctx, studentID, mock.MatchedBy(func(s *student_entity.Student) bool {
		return s.PersonalInfo.FullName == newFullName && s.PersonalInfo.Email == "john@example.com"
	})).Return(existingStudent, nil)

	// Execute
	updatedStudent, err := usecase.Update(ctx, studentID, updateDto)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, updatedStudent)
	assert.Equal(t, newFullName, updatedStudent.PersonalInfo.FullName)
	assert.Equal(t, "john@example.com", updatedStudent.PersonalInfo.Email)

	mockRepo.AssertExpectations(t)
}
