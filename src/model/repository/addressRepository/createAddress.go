package addressrepository

import (
	"context"
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
)

func (ar *addressRepository) CreateAddress(addressDomain domain.AddressDomainInterface) (domain.AddressDomainInterface, *resterr.RestErr) {

	addressEntity := converter.ConvertAddressDomainToEntity(addressDomain)

	addressID, err := ar.queries.CreateAddress(context.Background(), repository.CreateAddressParams{
		Street:       addressEntity.Street,
		Neighborhood: addressEntity.Neighborhood,
		State:        addressEntity.State,
		City:         addressEntity.City,
		ZipCode:      addressEntity.ZipCode,
		Notes:        addressEntity.Notes,
	})

	if err != nil {
		fmt.Print("err", err)
		return nil, resterr.InternalServerError("Internal server error", err.Error())

	}

	addressEntity.ID = addressID

	return converter.ConvertAddressEntityToDomain(*addressEntity), nil
}
