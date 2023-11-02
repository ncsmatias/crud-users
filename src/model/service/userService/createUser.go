package userservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[model] create user")

	logger.Info("new user created", journey)
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create user db")
	}
	return userDomainRepository, nil
}

func (ud *userDomainService) CreateUserTypeProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr) {

	journey := zap.String("journey", "[model] create user type professor")

	logger.Info("new user type professor created", journey)

	professorDomainRepository, err := ud.userRepository.CreateUserTypeProfessor(professorDomain)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create user db")
	}

	return professorDomainRepository, nil
}

func (ud *userDomainService) CreateUserTypeStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr) {

	journey := zap.String("journey", "[model] create user type student")

	logger.Info("new user type student created", journey)

	studentDomainRepository, err := ud.userRepository.CreateUserTypeStudent(studentDomain)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create user db")
	}

	return studentDomainRepository, nil
}
