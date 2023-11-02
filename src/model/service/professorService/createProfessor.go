package professorservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func (ps *ProfessorDomainService) CreateProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr) {

	professorDomainRepository, err := ps.professorRepository.CreateProfessor(professorDomain)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create professor db")
	}

	return professorDomainRepository, nil
}
