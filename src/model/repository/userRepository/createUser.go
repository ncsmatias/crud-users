package userrepository

import (
	"context"

	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (sr *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[repository] create user")
	value := converter.ConvertUserDomainToEntity(userDomain)

	userID, err := sr.queries.CreateUser(context.Background(),
		repository.CreateUserParams{
			Email:         value.Email,
			Name:          value.Name,
			Password:      value.Password,
			Role:          value.Role,
			Admin:         value.Admin,
			InstitutionID: value.InstitutionID,
		},
	)

	if err != nil {

		logger.Error("Error to create user", err, journey)
		return nil, resterr.InternalServerError("Internal sever error to execute query to create user", err.Error())
	}

	value.ID = userID

	logger.Info("New user created", journey)
	return converter.ConvertUserEntityToDomain(*value), nil
}

func (sr *userRepository) CreateUserTypeProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[repository] create user type professor")
	value := converter.ConvertProfessorDomainToEntity(professorDomain)

	professorID, err := sr.queries.CreateProfessor(context.Background(), repository.CreateProfessorParams{
		Department: value.Department,
		UserID:     value.UserID,
	})

	if err != nil {

		logger.Error("Error to create user type professor", err, journey)
		return nil, resterr.InternalServerError("Internal sever error to execute query to create professor", err.Error())
	}

	value.ID = professorID

	logger.Info("New user type professor created", journey)
	return converter.ConvertProfessorEntityToDomain(*value), nil
}

func (sr *userRepository) CreateUserTypeStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[repository] create user type student")
	value := converter.ConvertStudentDomainToEntity(studentDomain)

	studentID, err := sr.queries.CreateStudent(context.Background(), repository.CreateStudentParams{
		Course:      value.Course,
		StudentType: value.StudentType,
		ProfessorID: value.ProfessorID,
		UserID:      value.UserID,
	})

	if err != nil {

		logger.Error("Error to create user type student", err, journey)
		return nil, resterr.InternalServerError("Internal sever error to execute query to create student", err.Error())
	}

	value.ID = studentID

	logger.Info("New user type student created", journey)
	return converter.ConvertStudentEntityToDomain(*value), nil
}
