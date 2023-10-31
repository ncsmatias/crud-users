package converter

import (
	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) domain.UserDomainInterface {

	var name string
	if entity.Name.Valid {
		name = entity.Name.String
	} else {
		name = ""
	}

	var institutionID uuid.UUID

	if entity.InstitutionID.Valid {
		institutionID = entity.InstitutionID.UUID
	} else {

		institutionID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
	}

	userDomain := domain.NewUserDomain(
		entity.Email,
		name,
		entity.Password,
		entity.Role,
		entity.Admin,
		institutionID,
	)

	userDomain.SetID(entity.ID)

	return userDomain
}
