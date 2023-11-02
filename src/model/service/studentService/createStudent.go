package studentservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (sd *StudentDomainService) CreateStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[model] create student")

	studentDomainRepository, err := sd.studentRepository.CreateStudent(studentDomain)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create user db")
	}
	logger.Info("new user created", journey)
	return studentDomainRepository, nil
}
