package view

import (
	"github.com/ncsmatias/crud-users/src/controller/model/response"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func ConvertAddressDomainToResponse(addressDomain domain.AddressDomainInterface) *response.AddressResponse {

	return &response.AddressResponse{
		ID:           addressDomain.GetID(),
		Street:       addressDomain.GetStreet(),
		Neighborhood: addressDomain.GetNeighborhood(),
		City:         addressDomain.GetCity(),
		State:        addressDomain.GetCity(),
		ZipCode:      addressDomain.GetZipCode(),
		Notes:        addressDomain.GetNotes(),
	}
}

func ConvertUserDomainToResponse(
	userDomain domain.UserDomainInterface,
	professorDomain domain.ProfessorDomainInterface,
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

func ConvertInstitutionDomainToResponse(institutionDomain domain.InstitutionDomainInterface) response.InstitutionRequest {

	return response.InstitutionRequest{
		ID:              institutionDomain.GetID(),
		InstitutionType: institutionDomain.GetInstitutionType(),
		Name:            institutionDomain.GetName(),
		Phone:           institutionDomain.GetPhone(),
		AddressID:       institutionDomain.GetAddressID(),
	}
}
