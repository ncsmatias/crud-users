package converter

import (
	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/model/domain"
	"github.com/ncsmatias/crud-users/src/model/repository/entity"
)

func ConvertUserEntityToDomain(
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

func ConvertAddressEntityToDomain(
	addressEntity entity.AddressEntity,
) domain.AddressDomainInterface {

	var notes string
	if addressEntity.Notes.Valid {
		notes = addressEntity.Notes.String
	} else {
		notes = ""
	}

	addressDomain := domain.NewAddressDomain(
		addressEntity.Street,
		addressEntity.Neighborhood,
		addressEntity.City,
		addressEntity.State,
		addressEntity.ZipCode,
		notes,
	)

	addressDomain.SetID(addressEntity.ID)

	return addressDomain
}

func ConvertInstitutionEntityToDomain(
	institutionEntity entity.InstitutionEntity,
) domain.InstitutionDomainInterface {

	var phone string
	if institutionEntity.Phone.Valid {
		phone = institutionEntity.Phone.String
	} else {
		phone = ""
	}

	var addressID uuid.UUID
	if institutionEntity.AddressID.Valid {
		addressID = institutionEntity.AddressID.UUID
	} else {
		addressID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
	}

	institutionDomain := domain.NewInstitutionDomain(
		institutionEntity.InstitutionType,
		institutionEntity.Name,
		phone,
		addressID,
	)

	institutionDomain.SetID(institutionEntity.ID)

	return institutionDomain
}

func ConvertProfessorEntityToDomain(
	professorEntity entity.ProfessorEntity,
) domain.ProfessorDomainInterface {

	var userID uuid.UUID
	if professorEntity.UserID.Valid {
		userID = professorEntity.UserID.UUID
	} else {
		userID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
	}

	professorDomain := domain.NewProfessorDomain(
		professorEntity.Department,
		userID,
	)

	professorDomain.SetID(professorEntity.ID)

	return professorDomain

}

func ConvertStudentEntityToDomain(
	studentEntity entity.StudentEntity,
) domain.StudentDomainInterface {

	var professorID uuid.UUID
	if studentEntity.ProfessorID.Valid {
		professorID = studentEntity.ProfessorID.UUID
	} else {
		professorID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
	}

	var userID uuid.UUID
	if studentEntity.UserID.Valid {
		userID = studentEntity.UserID.UUID
	} else {
		userID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
	}

	var studentType string
	if studentEntity.StudentType.Valid {
		studentType = studentEntity.StudentType.String
	} else {
		studentType = ""
	}

	var course string
	if studentEntity.Course.Valid {
		course = studentEntity.Course.String
	} else {
		course = ""
	}

	studentDomain := domain.NewStudentDomain(
		course,
		studentType,
		professorID,
		userID,
	)

	studentDomain.SetID(studentEntity.ID)

	return studentDomain
}
