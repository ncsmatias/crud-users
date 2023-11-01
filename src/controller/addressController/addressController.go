package addresscontroller

import (
	"github.com/gin-gonic/gin"
	addressservice "github.com/ncsmatias/crud-users/src/model/service/addressService"
)

func NewAddressController(service addressservice.AddressDomainServiceInterface) AddressControllerInterface {
	return &addressController{service: service}
}

type AddressControllerInterface interface {
	CreateAddress(c *gin.Context)
	UpdateAddress(c *gin.Context)
	FindAddressByZipCode(c *gin.Context)
	DeleteAddress(c *gin.Context)
}

type addressController struct {
	service addressservice.AddressDomainServiceInterface
}
