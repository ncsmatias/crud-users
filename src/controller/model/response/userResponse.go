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
	Department    string    `json:"department,omitempty"`

	Course      string `json:"course,omitempty"`
	TypeStudent string `json:"type_student,omitempty"`

	ProfessorID uuid.UUID `json:"professor_id,omitempty"`
}
