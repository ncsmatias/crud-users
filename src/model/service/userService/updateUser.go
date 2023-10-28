package userservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain domain.UserDomainInterface) *resterr.RestErr {
	return nil
}
