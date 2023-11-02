package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/validation"
	"github.com/ncsmatias/crud-users/src/controller/model/request"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface domain.UserDomainInterface
)

func (uc *userController) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	journey := zap.String("journey", "[controller] create user")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
		userRequest.Role,
		userRequest.Admin,
		userRequest.InstitutionID,
	)

	userResult, err := uc.service.CreateUser(userDomain)
	if err != nil {

		logger.Error("Error service create user", err, journey)
		c.JSON(err.Code, err)
		return
	}

	var professorReult domain.ProfessorDomainInterface
	if userRequest.Role == "professor" {
		professorDomain := domain.NewProfessorDomain(userRequest.Department, userResult.GetID())

		professorReult, err = uc.service.CreateUserTypeProfessor(professorDomain)

		if err != nil {

			logger.Error("Error service create professor", err, journey)
			c.JSON(err.Code, err)
			return
		}
	}

	var studentResult domain.StudentDomainInterface
	if userRequest.Role == "student" {
		studentDomain := domain.NewStudentDomain(userRequest.Course, userRequest.TypeStudent, userRequest.ProfessorID, userResult.GetID())

		studentResult, err = uc.service.CreateUserTypeStudent(studentDomain)

		if err != nil {

			logger.Error("Error service create student", err, journey)
			c.JSON(err.Code, err)
			return
		}
	}

	logger.Info("New user created", journey, zap.String("user", userResult.ToString()))
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userResult, professorReult, studentResult))

}
