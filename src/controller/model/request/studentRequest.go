package request

import "github.com/google/uuid"

type StudentRequest struct {
	Course      string    `json:"course" binding:"required_if=Role student"`
	TypeStudent string    `json:"type_student" binding:"required_if=Role student,oneof=undergraduate master doctoral"`
	ProfessorID uuid.UUID `json:"professor_id" binding:"required_if=Role student"`
}
