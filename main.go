package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ncsmatias/crud-users/src/configuration/database/postgresdb"
	"github.com/ncsmatias/crud-users/src/controller/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := postgresdb.NewPostgresDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, conn)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
