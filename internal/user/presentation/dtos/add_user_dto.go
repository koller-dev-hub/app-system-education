package dtos

type AddUserDto struct {
	Name     string `json:"name" binding:"required" validate:"required,min=2,max=100" example:"John"`
	Surname  string `json:"surname" binding:"required" validate:"required,min=2,max=100" example:"Doe"`
	Nickname string `json:"nickname" binding:"required" validate:"required,min=2,max=50" example:"johnd"`
	Age      int32  `json:"age" binding:"required" validate:"required,gte=1,lte=130" example:"30"`
	Email    string `json:"email" binding:"required" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" validate:"required,min=8" example:"strongPassword123"`
}
