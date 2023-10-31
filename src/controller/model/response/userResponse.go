package response

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Role          string    `json:"role"`
	Admin         bool      `json:"admin"`
	InstitutionID uuid.UUID `json:"institution_id"`
}
