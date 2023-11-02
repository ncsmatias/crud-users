package professorrepositoty

import (
	"context"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
)

func (pr *professorRepository) CreateProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr) {

	value := converter.ConvertProfessorDomainToEntity(professorDomain)

	professorID, err := pr.queries.CreateProfessor(context.Background(), repository.CreateProfessorParams{
		Department: value.Department,
		UserID:     value.UserID,
	})

	if err != nil {
		return nil, resterr.InternalServerError("internal sever error to execute query to create user")
	}

	value.ID = professorID

	return converter.ConvertProfessorEntityToDomain(*value), nil
}
