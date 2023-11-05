package institutioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/validation"
	"github.com/ncsmatias/crud-users/src/controller/model/request"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/view"
	"go.uber.org/zap"
)

func (ic *institutionController) CreateInstitution(c *gin.Context) {
	var requestInstitution request.InstitutionRequest
	journey := zap.String("journey", "[controller] create user")

	if err := c.ShouldBindJSON(&requestInstitution); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewInstitutionDomain(
		requestInstitution.InstitutionType,
		requestInstitution.Name,
		requestInstitution.Phone,
		requestInstitution.ZipCode,
		requestInstitution.Street,
		requestInstitution.Number,
		requestInstitution.Neighborhood,
		requestInstitution.City,
		requestInstitution.State,
		requestInstitution.UF,
		requestInstitution.Country,
		requestInstitution.CountryCode,
	)

	result, err := ic.service.CreateInstitution(domain)

	if err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info("New user created", journey, zap.String("user", result.ToString()))
	c.JSON(http.StatusOK, view.ConvertInstitutionDomainToResponse(result))
}
