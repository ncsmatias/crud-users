package studentrepository

import (
	"database/sql"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
)

func NewStudentRepository(database *sql.DB) StudentRepositoryInterface {
	return &studentRepository{
		databaseConnection: database,
		queries:            repository.New(database),
	}
}

type StudentRepositoryInterface interface {
	CreateStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr)
}

type studentRepository struct {
	databaseConnection *sql.DB
	queries            *repository.Queries
}
