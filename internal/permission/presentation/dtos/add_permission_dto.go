package permission_dtos

type AddPermissionDto struct {
	UserID      string   `json:"user_id" binding:"required"`
	Modules     []string `json:"modules" binding:"required"`
	Actions     []string `json:"actions" binding:"required"`
	Level       string   `json:"level" binding:"required"`
	Description string   `json:"description" binding:"required"`
}
