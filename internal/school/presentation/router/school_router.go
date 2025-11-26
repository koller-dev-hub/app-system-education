package school_router

import (
	"github.com/gin-gonic/gin"
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
	{
		schools.POST("", handler.CreateSchool)
		schools.GET("", handler.FindAllSchool)
		schools.GET("/:id", handler.FindByIdSchool)
		schools.PUT("/:id", handler.UpdateSchool)
		schools.DELETE("/:id", handler.DeleteSchool)
	}
}
