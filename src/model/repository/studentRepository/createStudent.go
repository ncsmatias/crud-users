package studentrepository

import (
	"context"
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
)

func (sr *studentRepository) CreateStudent(studentDomain domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr) {

	value := converter.ConvertStudentDomainToEntity(studentDomain)

	fmt.Println("value", value)
	studentID, err := sr.queries.CreateStudent(context.Background(), repository.CreateStudentParams{
		Course:      value.Course,
		StudentType: value.StudentType,
		ProfessorID: value.ProfessorID,
		UserID:      value.UserID,
	})

	if err != nil {
		return nil, resterr.InternalServerError("internal sever error to execute query to create user")
	}

	value.ID = studentID

	return converter.ConvertStudentEntityToDomain(*value), nil
}
