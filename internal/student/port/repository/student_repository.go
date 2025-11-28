package port_student_repository

import (
	"context"
	"errors"

	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
)

type StudentRepository interface {
	Save(ctx context.Context, s *student_entity.Student) (*student_entity.Student, error)
	FindAll(ctx context.Context) ([]*student_entity.Student, error)
	FindById(ctx context.Context, id string) (*student_entity.Student, error)
	Update(ctx context.Context, id string, s *student_entity.Student) (*student_entity.Student, error)
	Delete(ctx context.Context, id string) error
}

var ErrNotFound = errors.New("student not found")
