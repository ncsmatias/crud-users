package studentrepository

import (
	"context"

	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (sr *studentRepository) CreateStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[service] create user type student")
	value := converter.ConvertStudentDomainToEntity(studentDomain)

	studentID, err := sr.queries.CreateStudent(context.Background(), repository.CreateStudentParams{
		Course:      value.Course,
		StudentType: value.StudentType,
		ProfessorID: value.ProfessorID,
		UserID:      value.UserID,
	})

	if err != nil {
		return nil, resterr.InternalServerError("Internal server error", err.Error())
	}

	value.ID = studentID

	logger.Info("New student create user", journey)
	return converter.ConvertStudentEntityToDomain(*value), nil
}
