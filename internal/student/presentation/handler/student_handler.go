package student_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	student_mapper "github.com/williamkoller/system-education/internal/student/application/mapper"
	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
	port_student_handler "github.com/williamkoller/system-education/internal/student/port/handler"
	port_student_repository "github.com/williamkoller/system-education/internal/student/port/repository"
	port_student_usecase "github.com/williamkoller/system-education/internal/student/port/usecase"
	student_dtos "github.com/williamkoller/system-education/internal/student/presentation/dtos"
)

type StudentHandler struct {
	usecase port_student_usecase.StudentUsecase
}

func NewStudentHandler(usecase port_student_usecase.StudentUsecase) *StudentHandler {
	return &StudentHandler{usecase: usecase}
}

var _ port_student_handler.StudentHandler = &StudentHandler{}

func (s *StudentHandler) CreateStudent(c *gin.Context) {
	var input student_dtos.AddStudentDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	student, err := s.usecase.Create(c.Request.Context(), input)
	if err != nil {
		var validationErr *student_entity.ValidationError
		if errors.As(err, &validationErr) {
			c.Status(http.StatusBadRequest)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resp := student_mapper.ToStudentResponse(student)
	c.JSON(http.StatusCreated, resp)
}

func (s *StudentHandler) FindAll(c *gin.Context) {
	students, err := s.usecase.FindAll(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	resp := student_mapper.ToStudentResponses(students)
	c.JSON(http.StatusOK, resp)
}

func (s *StudentHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	student, err := s.usecase.FindById(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, port_student_repository.ErrNotFound) {
			c.Status(http.StatusNotFound)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	resp := student_mapper.ToStudentResponse(student)
	c.JSON(http.StatusOK, resp)
}

func (s *StudentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input student_dtos.UpdateStudentDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	student, err := s.usecase.Update(c.Request.Context(), id, input)
	if err != nil {
		if errors.Is(err, port_student_repository.ErrNotFound) {
			c.Status(http.StatusNotFound)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		var validationErr *student_entity.ValidationError
		if errors.As(err, &validationErr) {
			c.Status(http.StatusBadRequest)
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
		c.Status(http.StatusInternalServerError)
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	resp := student_mapper.ToStudentResponse(student)
	c.JSON(http.StatusOK, resp)
}

func (s *StudentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := s.usecase.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, port_student_repository.ErrNotFound) {
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
