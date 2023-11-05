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

func ConvertInstitutionDomainToEntity(
	institutionDomain domain.InstitutionDomainInterface,
) *entity.InstitutionEntity {
	return &entity.InstitutionEntity{
		InstitutionType: institutionDomain.GetInstitutionType(),
		Name:            institutionDomain.GetName(),
		Phone:           sql.NullString{String: institutionDomain.GetPhone(), Valid: true},
		ZipCode:         sql.NullString{String: institutionDomain.GetZipCode(), Valid: true},
		Street:          sql.NullString{String: institutionDomain.GetStreet(), Valid: true},
		Number:          sql.NullString{String: institutionDomain.GetNumber(), Valid: true},
		Neighborhood:    sql.NullString{String: institutionDomain.GetNeighborhood(), Valid: true},
		City:            sql.NullString{String: institutionDomain.GetCity(), Valid: true},
		State:           sql.NullString{String: institutionDomain.GetState(), Valid: true},
		UF:              sql.NullString{String: institutionDomain.GetUF(), Valid: true},
		Country:         sql.NullString{String: institutionDomain.GetCountry(), Valid: true},
		CountryCode:     sql.NullString{String: institutionDomain.GetCountryCode(), Valid: true},
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

func ConvertStudentDomainToEntity(
	studentDomain domain.StudentDomainInterface,
) *entity.StudentEntity {

	return &entity.StudentEntity{
		Course:      sql.NullString{String: studentDomain.GetCourse(), Valid: true},
		StudentType: sql.NullString{String: studentDomain.GetStudentType(), Valid: true},
		ProfessorID: uuid.NullUUID{UUID: studentDomain.GetProfessorID(), Valid: true},
		UserID:      uuid.NullUUID{UUID: studentDomain.GetUserID(), Valid: true},
	}
}
