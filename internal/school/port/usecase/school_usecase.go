package port_school_usecase

import (
	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
	school_dtos "github.com/williamkoller/system-education/internal/school/presentation/dtos"
)

type SchoolUseCase interface {
	Create(input school_dtos.AddSchoolDto) (*school_entity.School, error)
	FindAll() ([]*school_entity.School, error)
	FindById(id string) (*school_entity.School, error)
	Update(id string, update school_dtos.UpdateSchoolDto) (*school_entity.School, error)
	Delete(id string) error
}							