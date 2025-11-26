package port_school_handler

import "github.com/gin-gonic/gin"

type SchoolHandler interface {
	CreateSchool(c *gin.Context)
	FindAllSchool(c *gin.Context)
	FindByIdSchool(c *gin.Context)
	UpdateSchool(c *gin.Context)
	DeleteSchool(c *gin.Context)
}
