package port_school_usecase

import (
    "context"
    school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
    school_dtos "github.com/williamkoller/system-education/internal/school/presentation/dtos"
)

type SchoolUseCase interface {
    Create(ctx context.Context, input school_dtos.AddSchoolDto) (*school_entity.School, error)
    FindAll(ctx context.Context) ([]*school_entity.School, error)
    FindById(ctx context.Context, id string) (*school_entity.School, error)
    Update(ctx context.Context, id string, update school_dtos.UpdateSchoolDto) (*school_entity.School, error)
    Delete(ctx context.Context, id string) error
}
