package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ncsmatias/crud-users/src/controller/routes"
	usercontroller "github.com/ncsmatias/crud-users/src/controller/userController"
	userservice "github.com/ncsmatias/crud-users/src/model/service/userService"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	service := userservice.NewUserDomainService()
	userController := usercontroller.NewUserController(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
