package addressservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func (as *addressDomainService) CreateAddress(addressDomain domain.AddressDomainInterface) (domain.AddressDomainInterface, *resterr.RestErr) {

	addressDomainRepository, err := as.addressRepository.CreateAddress(addressDomain)

	if err != nil {
		return nil, resterr.BadRequestError("Faild create address db")
	}

	return addressDomainRepository, nil
}
