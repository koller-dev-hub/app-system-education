package port_auth_usecase

import "context"

type AuthUsecase interface {
    Login(ctx context.Context, email, password string) (string, error)
    Profile(ctx context.Context, email string) (string, error)
}
