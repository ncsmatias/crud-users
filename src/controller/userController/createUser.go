package usercontroller

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

var (
	UserDomainInterface domain.UserDomainInterface
)

func (uc *userController) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	journey := zap.String("journey", "[controller] create user")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewUserDomain(userRequest.Email, userRequest.Name, userRequest.Password, userRequest.Admin)

	if err := uc.service.CreateUser(domain); err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info("New user created", journey)
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(domain))

}
