package view

import (
	"github.com/ncsmatias/crud-users/src/controller/model/response"
	"github.com/ncsmatias/crud-users/src/model/domain"
)

func ConvertUserDomainToResponse(userDomain domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:            "",
		Email:         userDomain.GetEmail(),
		Name:          userDomain.GetName(),
		Role:          userDomain.GetRole(),
		Admin:         userDomain.GetAdmin(),
		InstitutionID: userDomain.GetinstitutionID(),
	}
}
