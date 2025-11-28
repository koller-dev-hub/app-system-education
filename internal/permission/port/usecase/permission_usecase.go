package port_permission_usecase

import (
    "context"
    permission_entity "github.com/williamkoller/system-education/internal/permission/domain/entity"
    permission_dtos "github.com/williamkoller/system-education/internal/permission/presentation/dtos"
)

type PermissionUsecase interface {
    Create(ctx context.Context, input permission_dtos.AddPermissionDto) (*permission_entity.Permission, error)
    FindAll(ctx context.Context) ([]*permission_entity.Permission, error)
    FindById(ctx context.Context, id string) (*permission_entity.Permission, error)
    Update(ctx context.Context, id string, input permission_dtos.UpdatePermissionDto) (*permission_entity.Permission, error)
    Delete(ctx context.Context, id string) error
    FindPermissionByUserID(ctx context.Context, userID string) ([]*permission_entity.Permission, error)
}
