package userrepository

import (
	"context"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
)

func (sr *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr) {

	userID, err := sr.queries.CreateUser(context.Background(), repository.CreateUserParams{})

	if err != nil {
		return nil, resterr.InternalServerError("internal sever error to execute query to create user")

	}

	userDomain.SetID(userID)
	return userDomain, nil
}
