package institutioncontroller

import (
	"github.com/gin-gonic/gin"
	institutionservice "github.com/ncsmatias/crud-users/src/model/service/institutionService"
)

type InstitutionControllerInterface interface {
	CreateInstitution(c *gin.Context)
	FindInstitutionByID(c *gin.Context)
	UpdateInstitution(c *gin.Context)
	DeleteInstitution(c *gin.Context)
}

func NewInstitutionController(service institutionservice.InstitutionDomainServiceInterface) InstitutionControllerInterface {
	return &institutionController{service: service}
}

type institutionController struct {
	service institutionservice.InstitutionDomainServiceInterface
}
