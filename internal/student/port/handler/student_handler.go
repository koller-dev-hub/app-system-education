package port_student_handler

import (
	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	CreateStudent(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
