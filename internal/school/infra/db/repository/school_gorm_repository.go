package school_repository

import (
	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
	school_model "github.com/williamkoller/system-education/internal/school/infra/db/model"
	port_school_repository "github.com/williamkoller/system-education/internal/school/port/repository"
	"gorm.io/gorm"
)

type SchoolGormRepository struct {
	db *gorm.DB
}

func NewSchoolGormRepository(db *gorm.DB) *SchoolGormRepository {
	return &SchoolGormRepository{db: db}
}

var _ port_school_repository.SchoolRepository = &SchoolGormRepository{}

func (r *SchoolGormRepository) Save(s *school_entity.School) (*school_entity.School, error) {
	model := school_model.FromEntity(s)
	if err := r.db.Create(&model).Error; err != nil {
		return nil, err
	}
	return school_model.ToEntity(model), nil
}

func (r *SchoolGormRepository) Update(id string, s *school_entity.School) (*school_entity.School, error) {
	model := school_model.FromEntity(s)
	result := r.db.Model(&school_model.School{}).Where("id = ?", id).Updates(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, port_school_repository.ErrNotFound
	}

	return school_model.ToEntity(model), nil
}

func (r *SchoolGormRepository) Delete(id string) error {
	result := r.db.Unscoped().Delete(&school_model.School{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return port_school_repository.ErrNotFound
	}
	return nil
}

func (r *SchoolGormRepository) FindAll() ([]*school_entity.School, error) {
	var schools []*school_entity.School
	models := school_model.FromEntities(schools)
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	return school_model.ToEntities(models), nil
}

func (r *SchoolGormRepository) FindById(id string) (*school_entity.School, error) {
	var school *school_entity.School
	model := school_model.FromEntity(school)

	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, port_school_repository.ErrNotFound
		}
		return nil, err
	}

	return school_model.ToEntity(model), nil
}
