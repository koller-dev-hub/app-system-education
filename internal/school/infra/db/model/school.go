package school_model

import (
	"time"

	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
	"gorm.io/gorm"
)

type School struct {
	ID          string `gorm:"primaryKey;type:uuid"`
	Name        string
	Code        string
	Address     string
	City        string
	State       string
	ZipCode     string
	Country     string
	PhoneNumber string
	Email       string
	IsActive    bool
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (School) TableName() string {
	return "schools"
}

func FromEntity(s *school_entity.School) *School {
	if s == nil {
		return nil
	}
	return &School{
		ID:          s.ID,
		Name:        s.Name,
		Code:        s.Code,
		Address:     s.Address,
		City:        s.City,
		State:       s.State,
		ZipCode:     s.ZipCode,
		Country:     s.Country,
		PhoneNumber: s.PhoneNumber,
		Email:       s.Email,
		IsActive:    s.IsActive,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func FromEntities(ss []*school_entity.School) []*School {
	models := make([]*School, 0, len(ss))
	for _, s := range ss {
		models = append(models, FromEntity(s))
	}
	return models
}

func ToEntity(s *School) *school_entity.School {
	if s == nil {
		return nil
	}
	return &school_entity.School{
		ID:          s.ID,
		Name:        s.Name,
		Code:        s.Code,
		Address:     s.Address,
		City:        s.City,
		State:       s.State,
		ZipCode:     s.ZipCode,
		Country:     s.Country,
		PhoneNumber: s.PhoneNumber,
		Email:       s.Email,
		IsActive:    s.IsActive,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func ToEntities(ss []*School) []*school_entity.School {
	entities := make([]*school_entity.School, 0, len(ss))
	for _, s := range ss {
		entities = append(entities, ToEntity(s))
	}
	return entities
}

