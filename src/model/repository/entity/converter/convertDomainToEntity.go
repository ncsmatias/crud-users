package converter

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain domain.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:         domain.GetEmail(),
		Name:          sql.NullString{String: domain.GetName(), Valid: true},
		Role:          domain.GetRole(),
		Admin:         domain.GetAdmin(),
		InstitutionID: uuid.NullUUID{UUID: domain.GetinstitutionID(), Valid: true},
	}
}
