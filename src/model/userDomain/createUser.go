package userdomain

import (
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/logger"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *resterr.RestErr {
	journey := zap.String("journey", "[model] create user")

	logger.Info("new user created", journey)
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
