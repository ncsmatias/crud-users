package professorcontroller

import (
	"github.com/gin-gonic/gin"
	professorservice "github.com/ncsmatias/crud-users/src/model/service/professorService"
)

func NewProfessorController(service professorservice.ProfessorDomainServiceInterface) ProfessorControllerInterface {
	return &professorController{service: service}
}

type ProfessorControllerInterface interface {
	UpdateProfessor(c *gin.Context)
}

type professorController struct {
	service professorservice.ProfessorDomainServiceInterface
}
