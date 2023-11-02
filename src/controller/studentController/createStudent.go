package studentcontroller

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

func (sc *studentController) CreateStudent(c *gin.Context) {
	var studentRequest request.StudentRequest
	journey := zap.String("journey", "[controller] create student")

	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		logger.Error("Error ShouldBindJSON", err, journey)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewStudentDomain(studentRequest.Course, studentRequest.TypeStudent, studentRequest.ProfessorID, studentRequest.UserID)
	result, err := sc.service.CreateStudent(domain)

	if err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info("New student created", journey, zap.String("user", result.ToString()))
	c.JSON(http.StatusOK, view.ConvertStudentDomainToResponse(result))
}
