package request

import (
	"github.com/google/uuid"
	globaltypes "github.com/ncsmatias/crud-users/src/configuration/globalTypes"
)

type UserRequest struct {
	Email         string               `json:"email" binding:"required,email"`
	Password      string               `json:"password" binding:"required,min=8,containsany=!@$#*"`
	Name          string               `json:"name" binding:"max=128"`
	Role          globaltypes.UserRole `json:"role" binding:"required"`
	Admin         bool                 `json:"admin" binding:"required"`
	InstitutionID uuid.UUID            `json:"institution_id" binding:"required"`
}
