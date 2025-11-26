package school_usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
	school_dtos "github.com/williamkoller/system-education/internal/school/presentation/dtos"
)

type MockSchoolRepository struct {
	mock.Mock
}

func (m *MockSchoolRepository) Save(s *school_entity.School) (*school_entity.School, error) {
	args := m.Called(s)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*school_entity.School), args.Error(1)
}

func (m *MockSchoolRepository) Update(id string, s *school_entity.School) (*school_entity.School, error) {
	args := m.Called(id, s)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*school_entity.School), args.Error(1)
}

func (m *MockSchoolRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSchoolRepository) FindAll() ([]*school_entity.School, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*school_entity.School), args.Error(1)
}

func (m *MockSchoolRepository) FindById(id string) (*school_entity.School, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*school_entity.School), args.Error(1)
}

func TestSchoolUseCase_Create(t *testing.T) {
	t.Run("should create school successfully", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		input := school_dtos.AddSchoolDto{
			Name:        "Test School",
			Code:        "TS001",
			Address:     "123 Test St",
			City:        "Test City",
			State:       "TS",
			ZipCode:     "12345",
			Country:     "Test Country",
			PhoneNumber: "1234567890",
			Email:       "test@school.com",
			IsActive:    true,
			Description: "A test school",
		}

		mockRepo.On("Save", mock.AnythingOfType("*school_entity.School")).Return(&school_entity.School{
			ID:          "123",
			Name:        input.Name,
			Code:        input.Code,
			Address:     input.Address,
			City:        input.City,
			State:       input.State,
			ZipCode:     input.ZipCode,
			Country:     input.Country,
			PhoneNumber: input.PhoneNumber,
			Email:       input.Email,
			IsActive:    input.IsActive,
			Description: input.Description,
		}, nil)

		school, err := usecase.Create(input)

		assert.NoError(t, err)
		assert.NotNil(t, school)
		assert.Equal(t, input.Name, school.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when validation fails", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		input := school_dtos.AddSchoolDto{
			Name: "", // Invalid: Name is required
		}

		school, err := usecase.Create(input)

		assert.Error(t, err)
		assert.Nil(t, school)
		mockRepo.AssertNotCalled(t, "Save")
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		input := school_dtos.AddSchoolDto{
			Name:        "Test School",
			Code:        "TS001",
			Address:     "123 Test St",
			City:        "Test City",
			State:       "TS",
			ZipCode:     "12345",
			Country:     "Test Country",
			PhoneNumber: "1234567890",
			Email:       "test@school.com",
			IsActive:    true,
			Description: "A test school",
		}

		mockRepo.On("Save", mock.AnythingOfType("*school_entity.School")).Return(nil, errors.New("db error"))

		school, err := usecase.Create(input)

		assert.Error(t, err)
		assert.Nil(t, school)
		mockRepo.AssertExpectations(t)
	})
}

func TestSchoolUseCase_FindAll(t *testing.T) {
	t.Run("should return all schools", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		expectedSchools := []*school_entity.School{
			{ID: "1", Name: "School 1"},
			{ID: "2", Name: "School 2"},
		}

		mockRepo.On("FindAll").Return(expectedSchools, nil)

		schools, err := usecase.FindAll()

		assert.NoError(t, err)
		assert.Equal(t, expectedSchools, schools)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		mockRepo.On("FindAll").Return(nil, errors.New("db error"))

		schools, err := usecase.FindAll()

		assert.Error(t, err)
		assert.Nil(t, schools)
		mockRepo.AssertExpectations(t)
	})
}

func TestSchoolUseCase_FindById(t *testing.T) {
	t.Run("should return school by id", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		expectedSchool := &school_entity.School{ID: "123", Name: "Test School"}

		mockRepo.On("FindById", "123").Return(expectedSchool, nil)

		school, err := usecase.FindById("123")

		assert.NoError(t, err)
		assert.Equal(t, expectedSchool, school)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		mockRepo.On("FindById", "123").Return(nil, errors.New("db error"))

		school, err := usecase.FindById("123")

		assert.Error(t, err)
		assert.Nil(t, school)
		mockRepo.AssertExpectations(t)
	})
}

func TestSchoolUseCase_Update(t *testing.T) {
	t.Run("should update school successfully", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		id := "123"
		existingSchool := &school_entity.School{
			ID:          id,
			Name:        "Old Name",
			Description: "Old Description",
		}

		updateDto := school_dtos.UpdateSchoolDto{
			Name:        "Updated School",
			Description: "Updated Description",
		}

		updatedSchool := &school_entity.School{
			ID:          id,
			Name:        updateDto.Name,
			Description: updateDto.Description,
		}

		mockRepo.On("FindById", id).Return(existingSchool, nil)
		mockRepo.On("Update", id, updatedSchool).Return(updatedSchool, nil)

		school, err := usecase.Update(id, updateDto)

		assert.NoError(t, err)
		assert.Equal(t, updatedSchool, school)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		id := "123"
		existingSchool := &school_entity.School{
			ID:   id,
			Name: "Old Name",
		}

		updateDto := school_dtos.UpdateSchoolDto{
			Name: "Updated School",
		}

		updatedSchool := &school_entity.School{
			ID:   id,
			Name: "Updated School",
		}

		mockRepo.On("FindById", id).Return(existingSchool, nil)
		mockRepo.On("Update", id, updatedSchool).Return(nil, errors.New("db error"))

		school, err := usecase.Update(id, updateDto)

		assert.Error(t, err)
		assert.Nil(t, school)
		mockRepo.AssertExpectations(t)
	})
}

func TestSchoolUseCase_Delete(t *testing.T) {
	t.Run("should delete school successfully", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		id := "123"

		mockRepo.On("Delete", id).Return(nil)

		err := usecase.Delete(id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockSchoolRepository)
		usecase := NewSchoolUseCase(mockRepo)

		id := "123"

		mockRepo.On("Delete", id).Return(errors.New("db error"))

		err := usecase.Delete(id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
