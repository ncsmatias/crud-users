package routes

import (
	"github.com/gin-gonic/gin"
	professorcontroller "github.com/ncsmatias/crud-users/src/controller/professorController"
	studentcontroller "github.com/ncsmatias/crud-users/src/controller/studentController"
	usercontroller "github.com/ncsmatias/crud-users/src/controller/userController"
	professorservice "github.com/ncsmatias/crud-users/src/model/service/professorService"
	studentservice "github.com/ncsmatias/crud-users/src/model/service/studentService"
	userservice "github.com/ncsmatias/crud-users/src/model/service/userService"
)

func InitRoutes(r *gin.RouterGroup) {

	userService := userservice.NewUserDomainService()
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
