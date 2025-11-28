package port_user_usecase

import (
    "context"
    user_entity "github.com/williamkoller/system-education/internal/user/domain/entity"
    "github.com/williamkoller/system-education/internal/user/presentation/dtos"
)

type UserUsecase interface {
    Create(ctx context.Context, input dtos.AddUserDto) (*user_entity.User, error)
    FindAll(ctx context.Context) ([]*user_entity.User, error)
    FindByID(ctx context.Context, id string) (*user_entity.User, error)
    Update(ctx context.Context, id string, input dtos.UpdateUserDto) (*user_entity.User, error)
    Delete(ctx context.Context, id string) error
}
