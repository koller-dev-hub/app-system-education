package port_school_repository

import (
	"errors"

	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
)

type SchoolRepository interface {
	Save(s *school_entity.School) (*school_entity.School, error)
	Update(id string, s *school_entity.School) (*school_entity.School, error)
	Delete(id string) error
	FindAll() ([]*school_entity.School, error)
	FindById(id string) (*school_entity.School, error)
}

var ErrNotFound = errors.New("school not found")
