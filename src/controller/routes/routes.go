package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ncsmatias/crud-users/src/controller/create"
	"github.com/ncsmatias/crud-users/src/controller/delete"
	"github.com/ncsmatias/crud-users/src/controller/find"
	"github.com/ncsmatias/crud-users/src/controller/update"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/id/:id", find.FindUserByID)
	r.GET("/user/email/:email", find.FindUserByEmail)
	r.POST("/user", create.CreateUser)
	r.PUT("/user/:id", update.UpdateUser)
	r.DELETE("/user/:id", delete.DeleteUser)
}
