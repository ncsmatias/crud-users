package userservice

import (
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr) {
	journey := zap.String("journey", "[model] create user")

	logger.Info("new user created", journey)
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	fmt.Println("create user service", userDomainRepository, err)
	if err != nil {
		return nil, resterr.BadRequestError("Faild create user db")
	}
	return userDomainRepository, nil
}
