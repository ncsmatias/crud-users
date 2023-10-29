package professorservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func NewUserDomainService() ProfessorDomainServiceInterface {
	return &ProfessorDomainService{}
}

type ProfessorDomainService struct {
}

type ProfessorDomainServiceInterface interface {
	CreateProfessor(domain.ProfessorDomainInterface) *resterr.RestErr
	UpdateProfessor(string, domain.ProfessorDomainInterface) *resterr.RestErr
}
