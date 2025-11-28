package user_repository

import (
	"context"

	userEntity "github.com/williamkoller/system-education/internal/user/domain/entity"
	user_model "github.com/williamkoller/system-education/internal/user/infra/db/model"
	portUserRepository "github.com/williamkoller/system-education/internal/user/port/repository"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

var _ portUserRepository.UserRepository = &UserGormRepository{}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(ctx context.Context, u *userEntity.User) (*userEntity.User, error) {
	model := user_model.FromEntity(u)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}

	return user_model.ToEntity(model), nil
}

func (r *UserGormRepository) FindByID(ctx context.Context, id string) (*userEntity.User, error) {
	var user *userEntity.User

	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserGormRepository) FindAll(ctx context.Context) ([]*userEntity.User, error) {
	var users []*userEntity.User

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserGormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Unscoped().Delete(&user_model.User{}, "id = ?", id).Error
}

func (r *UserGormRepository) FindByEmail(ctx context.Context, email string) (*userEntity.User, error) {
	var user *userEntity.User
	model := user_model.FromEntity(user)

	if err := r.db.WithContext(ctx).First(&model, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return user_model.ToEntity(model), nil
}

func (r *UserGormRepository) Update(ctx context.Context, id string, u *userEntity.User) (*userEntity.User, error) {
	model := user_model.FromEntity(u)

	if err := r.db.WithContext(ctx).Model(&user_model.User{}).
		Where("id = ?", id).
		Updates(&model).Error; err != nil {
		return nil, err
	}

	return user_model.ToEntity(model), nil
}
