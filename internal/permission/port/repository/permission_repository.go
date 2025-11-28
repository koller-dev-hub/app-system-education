package port_permission_repository

import (
    "context"
    permission_entity "github.com/williamkoller/system-education/internal/permission/domain/entity"
)

type PermissionRepository interface {
    Save(ctx context.Context, p *permission_entity.Permission) (*permission_entity.Permission, error)
    FindAll(ctx context.Context) ([]*permission_entity.Permission, error)
    FindPermissionByUserID(ctx context.Context, userID string) ([]*permission_entity.Permission, error)
    Update(ctx context.Context, id string, p *permission_entity.Permission) (*permission_entity.Permission, error)
    Delete(ctx context.Context, id string) error
    FindByID(ctx context.Context, id string) (*permission_entity.Permission, error)
}
