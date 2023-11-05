package view

import (
	"github.com/ncsmatias/crud-users/src/controller/model/response"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func ConvertUserDomainToResponse(
	userDomain domain.UserDomainInterface,
	professorDomain domain.ProfessorDomainInterface,
	studentDomain domain.StudentDomainInterface,
) response.UserResponse {

	if professorDomain != nil {
		return response.UserResponse{
			ID:            userDomain.GetID(),
			Email:         userDomain.GetEmail(),
			Name:          userDomain.GetName(),
			Role:          userDomain.GetRole(),
			Admin:         userDomain.GetAdmin(),
			InstitutionID: userDomain.GetinstitutionID(),

			Department:  professorDomain.GetDepartment(),
			ProfessorID: professorDomain.GetID(),
		}
	}

	if studentDomain != nil {
		return response.UserResponse{
			ID:            userDomain.GetID(),
			Email:         userDomain.GetEmail(),
			Name:          userDomain.GetName(),
			Role:          userDomain.GetRole(),
			Admin:         userDomain.GetAdmin(),
			InstitutionID: userDomain.GetinstitutionID(),

			Course:      studentDomain.GetCourse(),
			ProfessorID: studentDomain.GetProfessorID(),
			TypeStudent: studentDomain.GetStudentType(),
		}
	}

	return response.UserResponse{
		ID:            userDomain.GetID(),
		Email:         userDomain.GetEmail(),
		Name:          userDomain.GetName(),
		Role:          userDomain.GetRole(),
		Admin:         userDomain.GetAdmin(),
		InstitutionID: userDomain.GetinstitutionID(),
	}
}

func ConvertProfessorDomainToResponse(professorDomain domain.ProfessorDomainInterface) response.ProfessorResponse {
	return response.ProfessorResponse{
		ID:         professorDomain.GetID(),
		Department: professorDomain.GetDepartment(),
		UserID:     professorDomain.GetUserID(),
	}
}

func ConvertStudentDomainToResponse(studentDomain domain.StudentDomainInterface) response.StudentResponse {

	return response.StudentResponse{
		ID:          studentDomain.GetID(),
		Course:      studentDomain.GetCourse(),
		TypeStudent: studentDomain.GetStudentType(),
		ProfessorID: studentDomain.GetProfessorID(),
	}
}

func ConvertInstitutionDomainToResponse(institutionDomain domain.InstitutionDomainInterface) response.InstitutionResponse {

	return response.InstitutionResponse{
		ID:              institutionDomain.GetID(),
		InstitutionType: institutionDomain.GetInstitutionType(),
		Name:            institutionDomain.GetName(),
		Phone:           institutionDomain.GetPhone(),
		ZipCode:         institutionDomain.GetZipCode(),
		Street:          institutionDomain.GetStreet(),
		Number:          institutionDomain.GetNumber(),
		Neighborhood:    institutionDomain.GetNeighborhood(),
		City:            institutionDomain.GetCity(),
		State:           institutionDomain.GetState(),
		UF:              institutionDomain.GetUF(),
		Country:         institutionDomain.GetCountry(),
		CountryCode:     institutionDomain.GetCountry(),
	}
}
