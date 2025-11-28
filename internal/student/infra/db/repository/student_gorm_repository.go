package student_repository

import (
	"context"
	"errors"

	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
	student_model "github.com/williamkoller/system-education/internal/student/infra/db/model"
	port_student_repository "github.com/williamkoller/system-education/internal/student/port/repository"
	"gorm.io/gorm"
)

type StudentGormRepository struct {
	db *gorm.DB
}

var _ port_student_repository.StudentRepository = &StudentGormRepository{}

func NewStudentGormRepository(db *gorm.DB) *StudentGormRepository {
	return &StudentGormRepository{db: db}
}

func (r *StudentGormRepository) Save(ctx context.Context, s *student_entity.Student) (*student_entity.Student, error) {
	model := student_model.FromEntity(s)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StudentGormRepository) FindAll(ctx context.Context) ([]*student_entity.Student, error) {
	var models []*student_model.Student
	if err := r.db.WithContext(ctx).Preload("School").Find(&models).Error; err != nil {
		return nil, err
	}
	return student_model.ToEntities(models), nil
}

func (r *StudentGormRepository) FindById(ctx context.Context, id string) (*student_entity.Student, error) {
	var model student_model.Student
	if err := r.db.WithContext(ctx).Preload("School").First(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port_student_repository.ErrNotFound
		}
		return nil, err
	}
	return student_model.ToEntity(&model), nil
}

func (r *StudentGormRepository) Update(ctx context.Context, id string, s *student_entity.Student) (*student_entity.Student, error) {
	var count int64
	r.db.WithContext(ctx).Model(&student_model.Student{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return nil, port_student_repository.ErrNotFound
	}

	model := student_model.FromEntity(s)
	model.ID = id
	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StudentGormRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&student_model.Student{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return port_student_repository.ErrNotFound
	}
	return nil
}
