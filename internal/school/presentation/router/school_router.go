package school_router

import (
	"github.com/gin-gonic/gin"
	permission_middleware "github.com/williamkoller/system-education/internal/permission/presentation/middleware"
	school_usecase "github.com/williamkoller/system-education/internal/school/application/usecase"
	school_repository "github.com/williamkoller/system-education/internal/school/infra/db/repository"
	school_handler "github.com/williamkoller/system-education/internal/school/presentation/handler"
	"gorm.io/gorm"
)

func SchoolRouter(g *gin.Engine, db *gorm.DB) {
	schools := g.Group("/schools")
	repo := school_repository.NewSchoolGormRepository(db)
	usecase := school_usecase.NewSchoolUseCase(repo)
	handler := school_handler.NewSchoolHandler(usecase)
	middleware := permission_middleware.NewPermissionMiddleware()
	{
		schools.POST("", middleware.ModuleAccessMiddleware([]string{"school"}, []string{"create"}), handler.CreateSchool)
		schools.GET("", middleware.ModuleAccessMiddleware([]string{"school"}, []string{"read"}), handler.FindAllSchool)
		schools.GET("/:id", middleware.ModuleAccessMiddleware([]string{"school"}, []string{"read"}), handler.FindByIdSchool)
		schools.PUT("/:id", middleware.ModuleAccessMiddleware([]string{"school"}, []string{"update"}), handler.UpdateSchool)
		schools.DELETE("/:id", middleware.ModuleAccessMiddleware([]string{"school"}, []string{"delete"}), handler.DeleteSchool)
	}
}
