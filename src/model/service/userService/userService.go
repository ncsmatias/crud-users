package userservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func NewUserDomainService() UserDomainServiceInterface {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainServiceInterface interface {
	CreateUser(domain.UserDomainInterface) *resterr.RestErr
	UpdateUser(string, domain.UserDomainInterface) *resterr.RestErr
	FindUser(string) (*domain.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
