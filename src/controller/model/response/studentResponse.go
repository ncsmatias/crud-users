package response

import "github.com/google/uuid"

type StudentResponse struct {
	ID          uuid.UUID `json:"id"`
	Course      string    `json:"course"`
	TypeStudent string    `json:"type_student"`
	ProfessorID uuid.UUID `json:"professor_id"`
}
