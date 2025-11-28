package permission_usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	permission_entity "github.com/williamkoller/system-education/internal/permission/domain/entity"
	permission_dtos "github.com/williamkoller/system-education/internal/permission/presentation/dtos"
)

type MockPermissionRepository struct {
	mock.Mock
}

func (m *MockPermissionRepository) Save(ctx context.Context, p *permission_entity.Permission) (*permission_entity.Permission, error) {
	args := m.Called(ctx, p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*permission_entity.Permission), args.Error(1)
}

func (m *MockPermissionRepository) FindAll(ctx context.Context) ([]*permission_entity.Permission, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*permission_entity.Permission), args.Error(1)
}

func (m *MockPermissionRepository) FindPermissionByUserID(ctx context.Context, userID string) ([]*permission_entity.Permission, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*permission_entity.Permission), args.Error(1)
}

func (m *MockPermissionRepository) Update(ctx context.Context, id string, p *permission_entity.Permission) (*permission_entity.Permission, error) {
	args := m.Called(ctx, id, p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*permission_entity.Permission), args.Error(1)
}

func (m *MockPermissionRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPermissionRepository) FindByID(ctx context.Context, id string) (*permission_entity.Permission, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*permission_entity.Permission), args.Error(1)
}

func TestPermissionUsecase_Create(t *testing.T) {
	t.Run("should create permission successfully", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		input := permission_dtos.AddPermissionDto{
			UserID:      "user-1",
			Modules:     []string{"module1"},
			Actions:     []string{"read"},
			Level:       "admin",
			Description: "test permission",
		}

		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*permission_entity.Permission")).Return(&permission_entity.Permission{
			ID:          "123",
			UserID:      input.UserID,
			Modules:     input.Modules,
			Actions:     input.Actions,
			Level:       input.Level,
			Description: input.Description,
		}, nil)

		permission, err := usecase.Create(context.Background(), input)

		assert.NoError(t, err)
		assert.NotNil(t, permission)
		assert.Equal(t, input.UserID, permission.UserID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when validation fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		input := permission_dtos.AddPermissionDto{
			UserID: "", // Invalid
		}

		permission, err := usecase.Create(context.Background(), input)

		assert.Error(t, err)
		assert.Nil(t, permission)
		mockRepo.AssertNotCalled(t, "Save")
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		input := permission_dtos.AddPermissionDto{
			UserID:      "user-1",
			Modules:     []string{"module1"},
			Actions:     []string{"read"},
			Level:       "admin",
			Description: "test permission",
		}

		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*permission_entity.Permission")).Return(nil, errors.New("db error"))

		permission, err := usecase.Create(context.Background(), input)

		assert.Error(t, err)
		assert.Nil(t, permission)
		mockRepo.AssertExpectations(t)
	})
}

func TestPermissionUsecase_FindAll(t *testing.T) {
	t.Run("should return all permissions", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		expectedPermissions := []*permission_entity.Permission{
			{ID: "1", UserID: "user-1"},
			{ID: "2", UserID: "user-2"},
		}

		mockRepo.On("FindAll", mock.Anything).Return(expectedPermissions, nil)

		permissions, err := usecase.FindAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedPermissions, permissions)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		mockRepo.On("FindAll", mock.Anything).Return(nil, errors.New("db error"))

		permissions, err := usecase.FindAll(context.Background())

		assert.Error(t, err)
		assert.Nil(t, permissions)
		mockRepo.AssertExpectations(t)
	})
}

func TestPermissionUsecase_FindById(t *testing.T) {
	t.Run("should return permission by id", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		expectedPermission := &permission_entity.Permission{ID: "123", UserID: "user-1"}

		mockRepo.On("FindByID", mock.Anything, "123").Return(expectedPermission, nil)

		permission, err := usecase.FindById(context.Background(), "123")

		assert.NoError(t, err)
		assert.Equal(t, expectedPermission, permission)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		mockRepo.On("FindByID", mock.Anything, "123").Return(nil, errors.New("db error"))

		permission, err := usecase.FindById(context.Background(), "123")

		assert.Error(t, err)
		assert.Nil(t, permission)
		mockRepo.AssertExpectations(t)
	})
}

func TestPermissionUsecase_Update(t *testing.T) {
	t.Run("should update permission successfully", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"
		modules := []string{"module2"}
		actions := []string{"write"}
		level := "user"
		description := "updated permission"

		input := permission_dtos.UpdatePermissionDto{
			Modules:     &modules,
			Actions:     &actions,
			Level:       &level,
			Description: &description,
		}

		existingPermission := &permission_entity.Permission{
			ID:          id,
			UserID:      "user-1",
			Modules:     []string{"module1"},
			Actions:     []string{"read"},
			Level:       "admin",
			Description: "test permission",
		}

		// Expected updated permission
		updatedPermission := &permission_entity.Permission{
			ID:          id,
			UserID:      "user-1",
			Modules:     *input.Modules,
			Actions:     *input.Actions,
			Level:       *input.Level,
			Description: *input.Description,
		}

		mockRepo.On("FindByID", mock.Anything, id).Return(existingPermission, nil)
		// We can't easily match the exact object pointer because it's modified in place,
		// but we can match the content or just use mock.AnythingOfType
		mockRepo.On("Update", mock.Anything, id, mock.AnythingOfType("*permission_entity.Permission")).Return(updatedPermission, nil)

		permission, err := usecase.Update(context.Background(), id, input)

		assert.NoError(t, err)
		assert.Equal(t, updatedPermission, permission)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when find by id fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"
		input := permission_dtos.UpdatePermissionDto{}

		mockRepo.On("FindByID", mock.Anything, id).Return(nil, permission_entity.ErrNotFound)

		permission, err := usecase.Update(context.Background(), id, input)

		assert.Error(t, err)
		assert.ErrorIs(t, err, permission_entity.ErrNotFound)
		assert.Nil(t, permission)
		mockRepo.AssertNotCalled(t, "Update")
	})

	t.Run("should return error when update fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"
		modules := []string{"module2"}
		input := permission_dtos.UpdatePermissionDto{
			Modules: &modules,
		}

		existingPermission := &permission_entity.Permission{
			ID:      id,
			Modules: []string{"module1"},
			Level:   "user",
		}

		mockRepo.On("FindByID", mock.Anything, id).Return(existingPermission, nil)
		mockRepo.On("Update", mock.Anything, id, mock.AnythingOfType("*permission_entity.Permission")).Return(nil, errors.New("db error"))

		permission, err := usecase.Update(context.Background(), id, input)

		assert.Error(t, err)
		assert.Nil(t, permission)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when validation fails during update", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"
		level := ""
		input := permission_dtos.UpdatePermissionDto{
			Level: &level,
		}

		existingPermission := &permission_entity.Permission{
			ID:    id,
			Level: "admin",
		}

		mockRepo.On("FindByID", mock.Anything, id).Return(existingPermission, nil)

		permission, err := usecase.Update(context.Background(), id, input)

		assert.Error(t, err)
		assert.Nil(t, permission)
		assert.Contains(t, err.Error(), "level cannot be empty")
		mockRepo.AssertNotCalled(t, "Update")
	})
}

func TestPermissionUsecase_Delete(t *testing.T) {
	t.Run("should delete permission successfully", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"

		// Mock FindByID first as Delete now calls it
		mockRepo.On("FindByID", mock.Anything, id).Return(&permission_entity.Permission{ID: id}, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(nil)

		err := usecase.Delete(context.Background(), id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when find by id fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"

		mockRepo.On("FindByID", mock.Anything, id).Return(nil, permission_entity.ErrNotFound)

		err := usecase.Delete(context.Background(), id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, permission_entity.ErrNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when delete fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		id := "123"

		mockRepo.On("FindByID", mock.Anything, id).Return(&permission_entity.Permission{ID: id}, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(errors.New("db error"))

		err := usecase.Delete(context.Background(), id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestPermissionUsecase_FindPermissionByUserID(t *testing.T) {
	t.Run("should return permissions by user id", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		userID := "user-1"
		expectedPermissions := []*permission_entity.Permission{
			{ID: "1", UserID: userID},
		}

		mockRepo.On("FindPermissionByUserID", mock.Anything, userID).Return(expectedPermissions, nil)

		permissions, err := usecase.FindPermissionByUserID(context.Background(), userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedPermissions, permissions)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockPermissionRepository)
		usecase := NewPermissionUsecase(mockRepo)

		userID := "user-1"

		mockRepo.On("FindPermissionByUserID", mock.Anything, userID).Return(nil, errors.New("db error"))

		permissions, err := usecase.FindPermissionByUserID(context.Background(), userID)

		assert.Error(t, err)
		assert.Nil(t, permissions)
		mockRepo.AssertExpectations(t)
	})
}
