package institutionrepository

import (
	"database/sql"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
)

func NewInstitutionRepository(database *sql.DB) InstitutionRepositoryInterface {
	return &institutionRepository{
		databaseConnection: database,
		queries:            repository.New(database),
	}
}

type InstitutionRepositoryInterface interface {
	CreateInstitution(institutionDomain domain.InstitutionDomainInterface) (domain.InstitutionDomainInterface, *resterr.RestErr)
}

type institutionRepository struct {
	databaseConnection *sql.DB
	queries            *repository.Queries
}
