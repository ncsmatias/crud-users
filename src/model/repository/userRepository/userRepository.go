package userrepository

import (
	"database/sql"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
)

func NewUserRepository(database *sql.DB) UserRepositoryInterface {
	return &userRepository{databaseConnection: database}
}

type UserRepositoryInterface interface {
	CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *resterr.RestErr)
}

type userRepository struct {
	databaseConnection *sql.DB
	queries            *repository.Queries
}
