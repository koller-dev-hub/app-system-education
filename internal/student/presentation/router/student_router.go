package student_router

import (
	"time"

	"github.com/gin-gonic/gin"
	infra_cryptography "github.com/williamkoller/system-education/internal/auth/infra/cryptography"
	auth_middleware "github.com/williamkoller/system-education/internal/auth/presentation/middleware"
	permission_middleware "github.com/williamkoller/system-education/internal/permission/presentation/middleware"
	student_usecase "github.com/williamkoller/system-education/internal/student/application/usecase"
	student_repository "github.com/williamkoller/system-education/internal/student/infra/db/repository"
	student_handler "github.com/williamkoller/system-education/internal/student/presentation/handler"
	"gorm.io/gorm"
)

func StudentRouter(g *gin.Engine, db *gorm.DB, secret string, expiresIn time.Duration) {
	studentGroup := g.Group("/students")
	repo := student_repository.NewStudentGormRepository(db)
	usecase := student_usecase.NewStudentUsecase(repo)
	handler := student_handler.NewStudentHandler(usecase)
	jwt := infra_cryptography.NewJWTTokenManager(secret, expiresIn)
	middleware := permission_middleware.NewPermissionMiddleware()
	{
		studentGroup.POST("/", auth_middleware.AuthMiddleware(jwt),
			middleware.ModuleAccessMiddleware([]string{"students"}, []string{"create"}),
			handler.CreateStudent)
		studentGroup.GET("/", auth_middleware.AuthMiddleware(jwt),
			middleware.ModuleAccessMiddleware([]string{"students"}, []string{"read"}),
			handler.FindAll)
		studentGroup.GET("/:id", auth_middleware.AuthMiddleware(jwt),
			middleware.ModuleAccessMiddleware([]string{"students"}, []string{"read"}),
			handler.FindById)
		studentGroup.PUT("/:id", auth_middleware.AuthMiddleware(jwt),
			middleware.ModuleAccessMiddleware([]string{"students"}, []string{"update"}),
			handler.Update)
		studentGroup.DELETE("/:id", auth_middleware.AuthMiddleware(jwt),
			middleware.ModuleAccessMiddleware([]string{"students"}, []string{"delete"}),
			handler.Delete)
	}
}
