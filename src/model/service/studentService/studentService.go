package studentservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func NewStudentDomainService() StudentDomainServiceInterface {
	return &StudentDomainService{}
}

type StudentDomainService struct {
}

type StudentDomainServiceInterface interface {
	CreateStudent(domain.StudentDomainInterface) *resterr.RestErr
	UpdateStudent(string, domain.StudentDomainInterface) *resterr.RestErr
}
