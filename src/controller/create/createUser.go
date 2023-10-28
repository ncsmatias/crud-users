package create

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/validation"
	"github.com/ncsmatias/crud-users/src/controller/model/request"
	"github.com/ncsmatias/crud-users/src/model/domain"
	userservice "github.com/ncsmatias/crud-users/src/model/service/userService"
	"go.uber.org/zap"
)

var (
	UserDomainInterface domain.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	journey := zap.String("journey", "[controller] create user")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewUserDomain(userRequest.Email, userRequest.Name, userRequest.Password, userRequest.Admin)
	service := userservice.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info("New user created", journey)
	fmt.Println(userRequest)

}
