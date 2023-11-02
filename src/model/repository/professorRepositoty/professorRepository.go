package professorrepositoty

import (
	"database/sql"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
)

func NewProfessoRepository(database *sql.DB) ProfessorRepositoryInterface {
	return &professorRepository{
		databaseConnection: database,
		queries:            repository.New(database),
	}
}

type ProfessorRepositoryInterface interface {
	CreateProfessor(professorDomain domain.ProfessorDomainInterface) (domain.ProfessorDomainInterface, *resterr.RestErr)
}

type professorRepository struct {
	databaseConnection *sql.DB
	queries            *repository.Queries
}
