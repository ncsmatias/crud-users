package create

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/validation"
	"github.com/ncsmatias/crud-users/src/controller/model/request"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	journey := zap.String("journey", "create user")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to create a new user", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("New user created", journey)
	fmt.Println(userRequest)

}
