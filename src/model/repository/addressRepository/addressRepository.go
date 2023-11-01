package addressrepository

import (
	"database/sql"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
)

func NewAddressRepository(database *sql.DB) AddressRepositoryInterface {
	return &addressRepository{
		databaseConnection: database,
		queries:            repository.New(database),
	}
}

type AddressRepositoryInterface interface {
	CreateAddress(addressDomain domain.AddressDomainInterface) (domain.AddressDomainInterface, *resterr.RestErr)
}

type addressRepository struct {
	databaseConnection *sql.DB
	queries            *repository.Queries
}
