package port_user_repository

import (
    "context"
    "errors"

    userEntity "github.com/williamkoller/system-education/internal/user/domain/entity"
)

type UserRepository interface {
    Save(ctx context.Context, u *userEntity.User) (*userEntity.User, error)
    FindByID(ctx context.Context, id string) (*userEntity.User, error)
    FindAll(ctx context.Context) ([]*userEntity.User, error)
    Delete(ctx context.Context, id string) error
    FindByEmail(ctx context.Context, email string) (*userEntity.User, error)
    Update(ctx context.Context, id string, u *userEntity.User) (*userEntity.User, error)
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)
