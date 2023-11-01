package addressservice

import (
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	addressrepository "github.com/ncsmatias/crud-users/src/model/repository/addressRepository"
)

func NewAddressDomainService(addressRepository addressrepository.AddressRepositoryInterface) AddressDomainServiceInterface {
	return &addressDomainService{addressRepository: addressRepository}
}

type addressDomainService struct {
	addressRepository addressrepository.AddressRepositoryInterface
}

type AddressDomainServiceInterface interface {
	CreateAddress(domain.AddressDomainInterface) (domain.AddressDomainInterface, *resterr.RestErr)
	UpdateAddress(string, domain.AddressDomainInterface) *resterr.RestErr
	FindAddressByZipCode(string) (domain.AddressDomainInterface, *resterr.RestErr)
	DeleteAddress(string) *resterr.RestErr
}
