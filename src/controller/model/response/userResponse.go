package response

import (
	"github.com/google/uuid"
	globaltypes "github.com/ncsmatias/crud-users/src/configuration/globalTypes"
)

type UserResponse struct {
	ID            uuid.UUID            `json:"id"`
	Email         string               `json:"email"`
	Name          string               `json:"name"`
	Role          globaltypes.UserRole `json:"role"`
	Admin         bool                 `json:"admin"`
	InstitutionID uuid.UUID            `json:"institution_id"`
}
