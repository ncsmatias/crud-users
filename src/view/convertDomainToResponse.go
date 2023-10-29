package view

import (
	"github.com/google/uuid"
	"github.com/ncsmatias/crud-users/src/controller/model/response"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func ConvertUserDomainToResponse(userDomain domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:            uuid.UUID{},
		Email:         userDomain.GetEmail(),
		Name:          userDomain.GetName(),
		Role:          userDomain.GetRole(),
		Admin:         userDomain.GetAdmin(),
		InstitutionID: userDomain.GetinstitutionID(),
	}
}

func ConvertProfessorDomainToResponse(professorDomain domain.ProfessorDomainInterface) response.ProfessorResponse {
	return response.ProfessorResponse{
		ID:         uuid.UUID{},
		Department: professorDomain.GetDepartment(),
		UserID:     professorDomain.GetUserID(),
	}
}

func ConvertStudentDomainToResponse(studentDomain domain.StudentDomainInterface) response.StudentResponse {

	return response.StudentResponse{
		ID:          uuid.UUID{},
		Course:      studentDomain.GetCourse(),
		TypeStudent: studentDomain.GetStudentType(),
		ProfessorID: studentDomain.GetProfessorID(),
	}
}
