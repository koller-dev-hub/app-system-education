package port_school_repository

import (
	"context"
	"errors"

	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
)

type SchoolRepository interface {
	Save(ctx context.Context, s *school_entity.School) (*school_entity.School, error)
	Update(ctx context.Context, id string, s *school_entity.School) (*school_entity.School, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]*school_entity.School, error)
	FindById(ctx context.Context, id string) (*school_entity.School, error)
}

var ErrNotFound = errors.New("school not found")
