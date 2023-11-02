package userrepository

import (
	"context"
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
)

func (sr *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr) {

	value := converter.ConvertUserDomainToEntity(userDomain)
	userID, err := sr.queries.CreateUser(context.Background(),
		repository.CreateUserParams{
			Email:         value.Email,
			Name:          value.Name,
			Password:      value.Password,
			Role:          value.Role,
			Admin:         value.Admin,
			InstitutionID: value.InstitutionID,
		},
	)

	if err != nil {
		fmt.Print("err", err)
		return nil, resterr.InternalServerError("internal sever error to execute query to create user")

	}

	value.ID = userID

	return converter.ConvertUserEntityToDomain(*value), nil
}

func (sr *userRepository) CreateUserTypeProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr) {

	value := converter.ConvertProfessorDomainToEntity(professorDomain)

	professorID, err := sr.queries.CreateProfessor(context.Background(), repository.CreateProfessorParams{
		Department: value.Department,
		UserID:     value.UserID,
	})

	if err != nil {
		return nil, resterr.InternalServerError("internal sever error to execute query to create user")
	}

	value.ID = professorID

	return converter.ConvertProfessorEntityToDomain(*value), nil
}
