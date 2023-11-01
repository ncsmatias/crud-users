package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	addresscontroller "github.com/ncsmatias/crud-users/src/controller/addressController"
	professorcontroller "github.com/ncsmatias/crud-users/src/controller/professorController"
	studentcontroller "github.com/ncsmatias/crud-users/src/controller/studentController"
	usercontroller "github.com/ncsmatias/crud-users/src/controller/userController"
	addressrepository "github.com/ncsmatias/crud-users/src/model/repository/addressRepository"
	userrepository "github.com/ncsmatias/crud-users/src/model/repository/userRepository"
	addressservice "github.com/ncsmatias/crud-users/src/model/service/addressService"
	professorservice "github.com/ncsmatias/crud-users/src/model/service/professorService"
	studentservice "github.com/ncsmatias/crud-users/src/model/service/studentService"
	userservice "github.com/ncsmatias/crud-users/src/model/service/userService"
)

func InitRoutes(r *gin.RouterGroup, conn *sql.DB) {

	addressRepository := addressrepository.NewAddressRepository(conn)
	addressService := addressservice.NewAddressDomainService(addressRepository)
	addressController := addresscontroller.NewAddressController(addressService)

	r.GET("/address/zip-code/:id", addressController.FindAddressByZipCode)
	r.POST("/address", addressController.CreateAddress)
	r.PUT("/address", addressController.UpdateAddress)
	r.DELETE("/address/id/:id", addressController.DeleteAddress)

	userRepository := userrepository.NewUserRepository(conn)

	userService := userservice.NewUserDomainService(userRepository)
	userController := usercontroller.NewUserController(userService)

	r.GET("/user/id/:id", userController.FindUserByID)
	r.GET("/user/email/:email", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	professorService := professorservice.NewProfessorDomainService()
	professorController := professorcontroller.NewProfessorController(professorService)

	r.POST("/professor", professorController.CreateProfessor)
	r.PUT("/professor/:id", professorController.UpdateProfessor)

	studentService := studentservice.NewStudentDomainService()
	studentController := studentcontroller.NewStudentController(studentService)

	r.POST("/student", studentController.CreateStudent)
	r.PUT("/student/:id", studentController.UpdateStudent)
}
