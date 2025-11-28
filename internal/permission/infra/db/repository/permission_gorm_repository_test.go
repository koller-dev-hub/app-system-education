package permission_repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	permission_entity "github.com/williamkoller/system-education/internal/permission/domain/entity"
	permission_model "github.com/williamkoller/system-education/internal/permission/infra/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PermissionGormRepositorySuite struct {
	suite.Suite
	db         *gorm.DB
	repository *PermissionGormRepository
}

func (s *PermissionGormRepositorySuite) SetupTest() {
	s.db = setupTestDB(s.T())
	s.repository = NewPermissionGormRepository(s.db)
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&permission_model.Permission{})
	assert.NoError(t, err)

	return db
}

func (s *PermissionGormRepositorySuite) TestCreate() {
	permission := &permission_entity.Permission{
		ID:          "perm-1",
		UserID:      "user-123",
		Modules:     []string{"module1"},
		Actions:     []string{"read"},
		Level:       "admin",
		Description: "test permission",
	}

	createdPermission, err := s.repository.Save(context.Background(), permission)

	s.NoError(err)
	s.NotNil(createdPermission)
	s.Equal(permission.ID, createdPermission.ID)
	s.Equal(permission.UserID, createdPermission.UserID)
	s.Equal(permission.Modules, createdPermission.Modules)
	s.Equal(permission.Actions, createdPermission.Actions)
	s.Equal(permission.Level, createdPermission.Level)
	s.Equal(permission.Description, createdPermission.Description)
}

func (s *PermissionGormRepositorySuite) TestCreate_Error() {
	permission := &permission_entity.Permission{
		ID:          "perm-2",
		UserID:      "user-123",
		Modules:     []string{"module1"},
		Actions:     []string{"read"},
		Level:       "admin",
		Description: "test permission",
	}

	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	createdPermission, err := s.repository.Save(context.Background(), permission)

	s.Error(err)
	s.Nil(createdPermission)
}

func (s *PermissionGormRepositorySuite) TestFindByID() {
	permission := &permission_entity.Permission{
		ID:          "perm-3",
		UserID:      "user-123",
		Modules:     []string{"module1"},
		Actions:     []string{"read"},
		Level:       "admin",
		Description: "test permission",
	}
	created, _ := s.repository.Save(context.Background(), permission)

	foundPermission, err := s.repository.FindByID(context.Background(), created.ID)

	s.NoError(err)
	s.NotNil(foundPermission)
	s.Equal(created.ID, foundPermission.ID)
	s.Equal(created.UserID, foundPermission.UserID)
}

func (s *PermissionGormRepositorySuite) TestFindByID_Error() {
	foundPermission, err := s.repository.FindByID(context.Background(), "some-id")

	s.Error(err)
	s.Equal(permission_entity.ErrNotFound, err)
	s.Nil(foundPermission)
}

func (s *PermissionGormRepositorySuite) TestFindByID_DBError() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	foundPermission, err := s.repository.FindByID(context.Background(), "some-id")

	s.Error(err)
	s.NotEqual(permission_entity.ErrNotFound, err)
	s.Nil(foundPermission)
}

func (s *PermissionGormRepositorySuite) TestFindAll() {
	permission1 := &permission_entity.Permission{ID: "perm-4", UserID: "user-1"}
	permission2 := &permission_entity.Permission{ID: "perm-5", UserID: "user-2"}
	s.repository.Save(context.Background(), permission1)
	s.repository.Save(context.Background(), permission2)

	permissions, err := s.repository.FindAll(context.Background())

	s.NoError(err)
	s.Len(permissions, 2)
}

func (s *PermissionGormRepositorySuite) TestFindAll_Error() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	permissions, err := s.repository.FindAll(context.Background())

	s.Error(err)
	s.Nil(permissions)
}

func (s *PermissionGormRepositorySuite) TestUpdate() {
	permission := &permission_entity.Permission{
		ID:          "perm-6",
		UserID:      "user-123",
		Modules:     []string{"module1"},
		Actions:     []string{"read"},
		Level:       "admin",
		Description: "test permission",
	}
	created, _ := s.repository.Save(context.Background(), permission)

	created.Description = "updated permission"
	updatedPermission, err := s.repository.Update(context.Background(), created.ID, created)

	s.NoError(err)
	s.NotNil(updatedPermission)
	s.Equal("updated permission", updatedPermission.Description)

	found, _ := s.repository.FindByID(context.Background(), created.ID)
	s.Equal("updated permission", found.Description)
}

func (s *PermissionGormRepositorySuite) TestUpdate_Error() {
	permission := &permission_entity.Permission{
		ID:     "perm-7",
		UserID: "user-123",
	}

	updatedPermission, err := s.repository.Update(context.Background(), permission.ID, permission)

	s.Error(err)
	s.Equal(permission_entity.ErrNotFound, err)
	s.Nil(updatedPermission)
}

func (s *PermissionGormRepositorySuite) TestUpdate_DBError() {
	permission := &permission_entity.Permission{
		ID:     "perm-7",
		UserID: "user-123",
	}
	created, _ := s.repository.Save(context.Background(), permission)

	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	updatedPermission, err := s.repository.Update(context.Background(), created.ID, created)

	s.Error(err)
	s.NotEqual(permission_entity.ErrNotFound, err)
	s.Nil(updatedPermission)
}

func (s *PermissionGormRepositorySuite) TestDelete() {
	permission := &permission_entity.Permission{ID: "perm-8", UserID: "user-123"}
	created, _ := s.repository.Save(context.Background(), permission)

	err := s.repository.Delete(context.Background(), created.ID)

	s.NoError(err)

	found, err := s.repository.FindByID(context.Background(), created.ID)
	s.Error(err)
	s.Nil(found)
}

func (s *PermissionGormRepositorySuite) TestDelete_Error() {
	err := s.repository.Delete(context.Background(), "some-id")

	s.Error(err)
	s.Equal(permission_entity.ErrNotFound, err)
}

func (s *PermissionGormRepositorySuite) TestDelete_DBError() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	err := s.repository.Delete(context.Background(), "some-id")

	s.Error(err)
	s.NotEqual(permission_entity.ErrNotFound, err)
}

func (s *PermissionGormRepositorySuite) TestFindPermissionByUserID() {
	permission1 := &permission_entity.Permission{ID: "perm-9", UserID: "user-123"}
	permission2 := &permission_entity.Permission{ID: "perm-10", UserID: "user-123"}
	permission3 := &permission_entity.Permission{ID: "perm-11", UserID: "user-456"}
	s.repository.Save(context.Background(), permission1)
	s.repository.Save(context.Background(), permission2)
	s.repository.Save(context.Background(), permission3)

	permissions, err := s.repository.FindPermissionByUserID(context.Background(), "user-123")

	s.NoError(err)
	s.Len(permissions, 2)
}

func (s *PermissionGormRepositorySuite) TestFindPermissionByUserID_Error() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()

	permissions, err := s.repository.FindPermissionByUserID(context.Background(), "user-123")

	s.Error(err)
	s.Nil(permissions)
}

func TestPermissionGormRepositorySuite(t *testing.T) {
	suite.Run(t, new(PermissionGormRepositorySuite))
}
