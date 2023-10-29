package professorservice

import (
	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func NewProfessorDomainService() ProfessorDomainServiceInterface {
	return &ProfessorDomainService{}
}

type ProfessorDomainService struct {
}

type ProfessorDomainServiceInterface interface {
	CreateProfessor(domain.ProfessorDomainInterface) *resterr.RestErr
	UpdateProfessor(uuid.UUID, domain.ProfessorDomainInterface) *resterr.RestErr
}
