package school_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	school_mapper "github.com/williamkoller/system-education/internal/school/application/mapper"
	port_school_handler "github.com/williamkoller/system-education/internal/school/port/handler"
	port_school_repository "github.com/williamkoller/system-education/internal/school/port/repository"
	port_school_usecase "github.com/williamkoller/system-education/internal/school/port/usecase"
	school_dtos "github.com/williamkoller/system-education/internal/school/presentation/dtos"
)

type SchoolHandler struct {
	usecase port_school_usecase.SchoolUseCase
}

var _ port_school_handler.SchoolHandler = &SchoolHandler{}

func NewSchoolHandler(usecase port_school_usecase.SchoolUseCase) *SchoolHandler {
	return &SchoolHandler{
		usecase: usecase,
	}
}

func (s *SchoolHandler) CreateSchool(c *gin.Context) {
	var input school_dtos.AddSchoolDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	school, err := s.usecase.Create(c.Request.Context(), input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resp := school_mapper.ToSchoolResponse(school)
	c.JSON(http.StatusCreated, resp)

}

func (s *SchoolHandler) FindAllSchool(c *gin.Context) {
	schools, err := s.usecase.FindAll(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	resp := school_mapper.ToSchoolResponses(schools)
	c.JSON(http.StatusOK, resp)
}

func (s *SchoolHandler) FindByIdSchool(c *gin.Context) {
	id := c.Param("id")
	school, err := s.usecase.FindById(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, port_school_repository.ErrNotFound) {
			c.Status(http.StatusNotFound)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	resp := school_mapper.ToSchoolResponse(school)
	c.JSON(http.StatusOK, resp)
}

func (s *SchoolHandler) UpdateSchool(c *gin.Context) {
	id := c.Param("id")
	var input school_dtos.UpdateSchoolDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	school, err := s.usecase.Update(c.Request.Context(), id, input)
	if err != nil {
		if errors.Is(err, port_school_repository.ErrNotFound) {
			c.Status(http.StatusNotFound)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	resp := school_mapper.ToSchoolResponse(school)
	c.JSON(http.StatusOK, resp)
}

func (s *SchoolHandler) DeleteSchool(c *gin.Context) {
	id := c.Param("id")
	err := s.usecase.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, port_school_repository.ErrNotFound) {
			c.Status(http.StatusNotFound)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.Status(http.StatusOK)
}
