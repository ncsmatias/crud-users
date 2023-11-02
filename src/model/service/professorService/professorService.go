package professorservice

import (
	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	professorrepositoty "github.com/ncsmatias/crud-users/src/model/repository/professorRepositoty"
)

func NewProfessorDomainService(professorRepository professorrepositoty.ProfessorRepositoryInterface) ProfessorDomainServiceInterface {
	return &ProfessorDomainService{professorRepository: professorRepository}
}

type ProfessorDomainService struct {
	professorRepository professorrepositoty.ProfessorRepositoryInterface
}

type ProfessorDomainServiceInterface interface {
	CreateProfessor(domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr)
	UpdateProfessor(uuid.UUID, domain.ProfessorDomainInterface) *resterr.RestErr
}
