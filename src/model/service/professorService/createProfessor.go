package professorservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (ps *ProfessorDomainService) CreateProfessor(professorDomain domain.ProfessorDomainInterface) *resterr.RestErr {
	journey := zap.String("journey", "[model] create professor")

	logger.Info("new user created", journey)
	return nil
}
