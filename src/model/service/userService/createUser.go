package userservice

import (
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) *resterr.RestErr {
	journey := zap.String("journey", "[model] create user")

	logger.Info("new user created", journey)
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetEmail())
	return nil
}
