package studentservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (sd *StudentDomainService) CreateStudent(studentDomain domain.StudentDomainInterface) *resterr.RestErr {
	journey := zap.String("journey", "[model] create student")

	logger.Info("new user created", journey)
	return nil
}
