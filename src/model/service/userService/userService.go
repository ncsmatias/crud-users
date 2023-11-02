package userservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	userrepository "github.com/ncsmatias/crud-users/src/model/repository/userRepository"
)

func NewUserDomainService(userRepository userrepository.UserRepositoryInterface) UserDomainServiceInterface {
	return &userDomainService{userRepository: userRepository}
}

type userDomainService struct {
	userRepository userrepository.UserRepositoryInterface
}

type UserDomainServiceInterface interface {
	CreateUser(domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr)
	CreateUserTypeProfessor(domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr)
	UpdateUser(string, domain.UserDomainInterface) *resterr.RestErr
	FindUser(string) (*domain.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
