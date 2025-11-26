package school_dtos

type AddSchoolDto struct {
	Name        string `json:"name" binding:"required" validate:"required"`
	Code        string `json:"code" binding:"required" validate:"required"`
	Address     string `json:"address" binding:"required" validate:"required"`
	City        string `json:"city" binding:"required" validate:"required"`
	State       string `json:"state" binding:"required" validate:"required"`
	ZipCode     string `json:"zip_code" binding:"required" validate:"required"`
	Country     string `json:"country" binding:"required" validate:"required"`
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required"`
	Email       string `json:"email" binding:"required" validate:"required"`
	IsActive    bool   `json:"is_active" binding:"required" validate:"required"`
	Description string `json:"description" binding:"required" validate:"required"`
}
