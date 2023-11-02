package studentservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	studentrepository "github.com/ncsmatias/crud-users/src/model/repository/studentRepository"
)

func NewStudentDomainService(studentRepository studentrepository.StudentRepositoryInterface) StudentDomainServiceInterface {
	return &StudentDomainService{studentRepository: studentRepository}
}

type StudentDomainService struct {
	studentRepository studentrepository.StudentRepositoryInterface
}

type StudentDomainServiceInterface interface {
	CreateStudent(domain.StudentDomainInterface) (domain.StudentDomainInterface, *resterr.RestErr)
	UpdateStudent(string, domain.StudentDomainInterface) *resterr.RestErr
}
