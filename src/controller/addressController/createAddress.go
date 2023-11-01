package addresscontroller

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

func (ac *addressController) CreateAddress(c *gin.Context) {
	var addressRequest request.AddressRequest
	journey := zap.String("journey", "[controller] create address")

	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewAddressDomain(
		addressRequest.Street,
		addressRequest.Neighborhood,
		addressRequest.City,
		addressRequest.State,
		addressRequest.ZipCode,
		addressRequest.Notes,
	)

	result, err := ac.service.CreateAddress(domain)

	if err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info("New address created", journey, zap.String("address", result.ToString()))
	c.JSON(http.StatusOK, view.ConvertAddressDomainToResponse(result))

	return
}
