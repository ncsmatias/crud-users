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
		AddressID:       value.AddressID,
	})

	if err != nil {
		fmt.Print("err", err)
		return nil, resterr.InternalServerError("internal sever error to execute query to create user")

	}

	value.ID = institutionID

	return converter.ConvertInstitutionEntityToDomain(*value), nil
}
