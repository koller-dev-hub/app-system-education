package school_router

import (
	"time"

	"github.com/gin-gonic/gin"
	infra_cryptography "github.com/williamkoller/system-education/internal/auth/infra/cryptography"
	auth_middleware "github.com/williamkoller/system-education/internal/auth/presentation/middleware"
	permission_middleware "github.com/williamkoller/system-education/internal/permission/presentation/middleware"
	school_usecase "github.com/williamkoller/system-education/internal/school/application/usecase"
	school_repository "github.com/williamkoller/system-education/internal/school/infra/db/repository"
	school_handler "github.com/williamkoller/system-education/internal/school/presentation/handler"
	"gorm.io/gorm"
)

func SchoolRouter(g *gin.Engine, db *gorm.DB, secret string, expiresIn time.Duration) {
	schools := g.Group("/schools")
	repo := school_repository.NewSchoolGormRepository(db)
	usecase := school_usecase.NewSchoolUseCase(repo)
	handler := school_handler.NewSchoolHandler(usecase)
	middleware := permission_middleware.NewPermissionMiddleware()
	jwt := infra_cryptography.NewJWTTokenManager(secret, expiresIn)

	{
		schools.POST("", auth_middleware.AuthMiddleware(jwt), middleware.ModuleAccessMiddleware([]string{"schools"}, []string{"create"}), handler.CreateSchool)
		schools.GET("", auth_middleware.AuthMiddleware(jwt), middleware.ModuleAccessMiddleware([]string{"schools"}, []string{"read"}), handler.FindAllSchool)
		schools.GET("/:id", auth_middleware.AuthMiddleware(jwt), middleware.ModuleAccessMiddleware([]string{"schools"}, []string{"read"}), handler.FindByIdSchool)
		schools.PUT("/:id", auth_middleware.AuthMiddleware(jwt), middleware.ModuleAccessMiddleware([]string{"schools"}, []string{"update"}), handler.UpdateSchool)
		schools.DELETE("/:id", auth_middleware.AuthMiddleware(jwt), middleware.ModuleAccessMiddleware([]string{"schools"}, []string{"delete"}), handler.DeleteSchool)
	}
}
