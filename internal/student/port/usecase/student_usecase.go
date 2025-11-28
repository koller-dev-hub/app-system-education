package port_student_usecase

import (
	"context"

	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
	student_dtos "github.com/williamkoller/system-education/internal/student/presentation/dtos"
)

type StudentUsecase interface {
	Create(ctx context.Context, input student_dtos.AddStudentDto) (*student_entity.Student, error)
	FindAll(ctx context.Context) ([]*student_entity.Student, error)
	FindById(ctx context.Context, id string) (*student_entity.Student, error)
	Update(ctx context.Context, id string, input student_dtos.UpdateStudentDto) (*student_entity.Student, error)
	Delete(ctx context.Context, id string) error
}
