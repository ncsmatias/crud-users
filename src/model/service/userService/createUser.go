package userservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr) {

	journey := zap.String("journey", "[service] create user")
	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {

		logger.Error("Error create repository user", err, journey)
		return nil, resterr.BadRequestError("Faild create user db")
	}

	logger.Info("New user created", journey)
	return userDomainRepository, nil
}

func (ud *userDomainService) CreateUserTypeProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr) {

	journey := zap.String("journey", "[service] create user type professor")
	professorDomainRepository, err := ud.userRepository.CreateUserTypeProfessor(professorDomain)

	if err != nil {

		logger.Error("Error create repository user type professor", err, journey)
		return nil, resterr.BadRequestError("Faild create user db")
	}

	logger.Info("New user type professor created", journey)
	return professorDomainRepository, nil
}

func (ud *userDomainService) CreateUserTypeStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr) {

	journey := zap.String("journey", "[service] create user type student")
	studentDomainRepository, err := ud.userRepository.CreateUserTypeStudent(studentDomain)

	if err != nil {

		logger.Error("Error create repository user type student", err, journey)
		return nil, resterr.BadRequestError("Faild create user db")
	}

	logger.Info("New user type student created", journey)
	return studentDomainRepository, nil
}
