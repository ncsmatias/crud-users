package usercontroller

import (
	"github.com/gin-gonic/gin"
	userservice "github.com/ncsmatias/crud-users/src/model/service/userService"
)

func NewUserController(service userservice.UserDomainServiceInterface) UserControllerInterface {
	return &userController{service: service}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	service userservice.UserDomainServiceInterface
}
