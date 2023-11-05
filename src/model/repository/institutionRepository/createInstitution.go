package institutionrepository

import (
	"context"
	"fmt"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository"
	"github.com/ncsmatias/crud-users/src/model/repository/entity/converter"
)

func (ir *institutionRepository) CreateInstitution(institutionDomain domain.InstitutionDomainInterface) (domain.InstitutionDomainInterface, *resterr.RestErr) {

	value := converter.ConvertInstitutionDomainToEntity(institutionDomain)

	institutionID, err := ir.queries.CreateInstitution(context.Background(), repository.CreateInstitutionParams{
		InstitutionType: value.InstitutionType,
		Name:            value.Name,
		Phone:           value.Phone,
		ZipCode:         value.ZipCode,
		Street:          value.Street,
		Number:          value.Number,
		Neighborhood:    value.Neighborhood,
		City:            value.City,
		State:           value.State,
		Uf:              value.UF,
		Country:         value.Country,
		CountryCode:     value.CountryCode,
	})

	if err != nil {
		fmt.Print("err", err)
		return nil, resterr.InternalServerError("Internal server error", err.Error())

	}

	value.ID = institutionID

	return converter.ConvertInstitutionEntityToDomain(*value), nil
}
