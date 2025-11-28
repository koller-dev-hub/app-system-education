package school_dtos

type UpdateSchoolDto struct {
	Name        *string `json:"name"`
	Code        *string `json:"code"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	ZipCode     *string `json:"zip_code"`
	Country     *string `json:"country"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	IsActive    *bool   `json:"is_active"`
	Description *string `json:"description"`
}
