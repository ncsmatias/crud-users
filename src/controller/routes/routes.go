package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/ncsmatias/crud-users/src/controller/userController"
)

func InitRoutes(r *gin.RouterGroup, userController usercontroller.UserControllerInterface) {

	r.GET("/user/id/:id", userController.FindUserByID)
	r.GET("/user/email/:email", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)
}
