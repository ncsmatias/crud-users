package studentcontroller

import (
	"github.com/gin-gonic/gin"
	studentservice "github.com/ncsmatias/crud-users/src/model/service/studentService"
)

func NewStudentController(service studentservice.StudentDomainServiceInterface) StudentControllerInterface {
	return &studentController{service: service}
}

type StudentControllerInterface interface {
	CreateStudent(c *gin.Context)
	UpdateStudent(c *gin.Context)
}

type studentController struct {
	service studentservice.StudentDomainServiceInterface
}
