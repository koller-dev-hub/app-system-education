package school_mapper

import (
	"time"

	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
)

type SchoolResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	ZipCode     string    `json:"zipCode"`
	Country     string    `json:"country"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	IsActive    bool      `json:"isActive"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToSchoolResponse(school *school_entity.School) *SchoolResponse {
	return &SchoolResponse{
		ID:          school.ID,
		Name:        school.Name,
		Code:        school.Code,
		Address:     school.Address,
		City:        school.City,
		State:       school.State,
		ZipCode:     school.ZipCode,
		Country:     school.Country,
		PhoneNumber: school.PhoneNumber,
		Email:       school.Email,
		IsActive:    school.IsActive,
		Description: school.Description,
		CreatedAt:   school.CreatedAt,
		UpdatedAt:   school.UpdatedAt,
	}
}

func ToSchoolResponses(schools []*school_entity.School) []*SchoolResponse {
	responses := make([]*SchoolResponse, 0, len(schools))
	for _, school := range schools {
		responses = append(responses, ToSchoolResponse(school))
	}
	return responses
}
