package response

import "github.com/google/uuid"

type ProfessorResponse struct {
	ID         uuid.UUID `json:"id"`
	Department string    `json:"department"`
	UserID     uuid.UUID `json:"user_id"`
}
