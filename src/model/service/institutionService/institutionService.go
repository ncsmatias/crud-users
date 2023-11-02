package institutionservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	institutionrepository "github.com/ncsmatias/crud-users/src/model/repository/institutionRepository"
)

type InstitutionDomainServiceInterface interface {
	CreateInstitution(domain.InstitutionDomainInterface) (domain.InstitutionDomainInterface, *resterr.RestErr)
	FindInstitutionByID(string) (domain.InstitutionDomainInterface, *resterr.RestErr)
	UpdateInstitution(string, domain.InstitutionDomainInterface) *resterr.RestErr
	DeleteInstitution(string) *resterr.RestErr
}

func NewInstitutionDomainService(institutionRepository institutionrepository.InstitutionRepositoryInterface) InstitutionDomainServiceInterface {
	return &institutionDomainService{repository: institutionRepository}
}

type institutionDomainService struct {
	repository institutionrepository.InstitutionRepositoryInterface
}
