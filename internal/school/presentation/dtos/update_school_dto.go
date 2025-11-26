package school_dtos

type UpdateSchoolDto struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zipCode"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	IsActive    bool   `json:"isActive"`
	Description string `json:"description"`
}
