package professorrepositoty

import (
	"database/sql"

	"github.com/ncsmatias/crud-users/src/model/repository"
)

func NewProfessoRepository(database *sql.DB) ProfessorRepositoryInterface {
	return &professorRepository{
		databaseConnection: database,
		queries:            repository.New(database),
	}
}

type ProfessorRepositoryInterface interface {
}

type professorRepository struct {
	databaseConnection *sql.DB
	queries            *repository.Queries
}
