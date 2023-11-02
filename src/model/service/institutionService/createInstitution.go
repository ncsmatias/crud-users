package institutionservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func (is *institutionDomainService) CreateInstitution(domainInstitution domain.InstitutionDomainInterface) (domain.InstitutionDomainInterface, *resterr.RestErr) {

	institutionDomainService, err := is.repository.CreateInstitution(domainInstitution)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create user db")
	}

	return institutionDomainService, nil
}
