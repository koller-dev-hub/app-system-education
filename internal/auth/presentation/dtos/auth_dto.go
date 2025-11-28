package auth_dtos

type AuthDto struct {
	Email    string `json:"email" binding:"required" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" validate:"required,min=8" example:"strongPassword123"`
}
