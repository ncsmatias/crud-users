package professorcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/validation"
	"github.com/ncsmatias/crud-users/src/controller/model/request"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/view"
	"go.uber.org/zap"
)

func (pc *professorController) CreateProfessor(c *gin.Context) {
	var professorRequest request.ProfessorRequest
	journey := zap.String("journey", "[controller] create professor")

	if err := c.ShouldBindJSON(&professorRequest); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println("id", professorRequest.UserID)

	domain := domain.NewProfessorDomain(professorRequest.Department, professorRequest.UserID)

	if err := pc.service.CreateProfessor(domain); err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info("New professor created", journey, zap.String("user", domain.ToString()))
	c.JSON(http.StatusOK, view.ConvertProfessorDomainToResponse(domain))
}
