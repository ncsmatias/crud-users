package converter

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository/entity"
)

func ConvertUserDomainToEntity(
	domain domain.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:         domain.GetEmail(),
		Name:          sql.NullString{String: domain.GetName(), Valid: true},
		Password:      domain.GetPassword(),
		Role:          domain.GetRole(),
		Admin:         domain.GetAdmin(),
		InstitutionID: uuid.NullUUID{UUID: domain.GetinstitutionID(), Valid: true},
	}
}

func ConvertAddressDomainToEntity(
	addressDomain domain.AddressDomainInterface,
) *entity.AddressEntity {
	return &entity.AddressEntity{
		Street:       addressDomain.GetStreet(),
		Neighborhood: addressDomain.GetNeighborhood(),
		City:         addressDomain.GetCity(),
		State:        addressDomain.GetState(),
		ZipCode:      addressDomain.GetZipCode(),
		Notes:        sql.NullString{String: addressDomain.GetNotes(), Valid: true},
	}
}

func ConvertInstitutionDomainToEntity(
	institutionDomain domain.InstitutionDomainInterface,
) *entity.InstitutionEntity {
	return &entity.InstitutionEntity{
		InstitutionType: institutionDomain.GetInstitutionType(),
		Name:            institutionDomain.GetName(),
		Phone:           sql.NullString{String: institutionDomain.GetPhone(), Valid: true},
		AddressID:       uuid.NullUUID{UUID: institutionDomain.GetAddressID(), Valid: true},
	}
}
func ConvertProfessorDomainToEntity(
	professorDomain domain.ProfessorDomainInterface,
) *entity.ProfessorEntity {
	return &entity.ProfessorEntity{
		Department: professorDomain.GetDepartment(),
		UserID:     uuid.NullUUID{UUID: professorDomain.GetUserID(), Valid: true},
	}
}
