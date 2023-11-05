package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	institutioncontroller "github.com/ncsmatias/crud-users/src/controller/institutionController"
	professorcontroller "github.com/ncsmatias/crud-users/src/controller/professorController"
	studentcontroller "github.com/ncsmatias/crud-users/src/controller/studentController"
	usercontroller "github.com/ncsmatias/crud-users/src/controller/userController"
	institutionrepository "github.com/ncsmatias/crud-users/src/model/repository/institutionRepository"
	professorrepositoty "github.com/ncsmatias/crud-users/src/model/repository/professorRepositoty"
	studentrepository "github.com/ncsmatias/crud-users/src/model/repository/studentRepository"
	userrepository "github.com/ncsmatias/crud-users/src/model/repository/userRepository"
	institutionservice "github.com/ncsmatias/crud-users/src/model/service/institutionService"
	professorservice "github.com/ncsmatias/crud-users/src/model/service/professorService"
	studentservice "github.com/ncsmatias/crud-users/src/model/service/studentService"
	userservice "github.com/ncsmatias/crud-users/src/model/service/userService"
)

func InitRoutes(r *gin.RouterGroup, conn *sql.DB) {

	institutionRepository := institutionrepository.NewInstitutionRepository(conn)
	institutionService := institutionservice.NewInstitutionDomainService(institutionRepository)
	institutionController := institutioncontroller.NewInstitutionController(institutionService)

	r.GET("/institution/id/:id", institutionController.FindInstitutionByID)
	r.POST("/institution/", institutionController.CreateInstitution)
	r.PUT("/institution/id/:id", institutionController.UpdateInstitution)
	r.DELETE("/institution/id/:id", institutionController.DeleteInstitution)

	userRepository := userrepository.NewUserRepository(conn)

	userService := userservice.NewUserDomainService(userRepository)
	userController := usercontroller.NewUserController(userService)

	r.GET("/user/id/:id", userController.FindUserByID)
	r.GET("/user/email/:email", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	professorRepository := professorrepositoty.NewProfessoRepository(conn)
	professorService := professorservice.NewProfessorDomainService(professorRepository)
	professorController := professorcontroller.NewProfessorController(professorService)

	r.PUT("/professor/:id", professorController.UpdateProfessor)

	studentRepository := studentrepository.NewStudentRepository(conn)
	studentService := studentservice.NewStudentDomainService(studentRepository)
	studentController := studentcontroller.NewStudentController(studentService)

	r.PUT("/student/:id", studentController.UpdateStudent)
}
