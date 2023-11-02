package request

import (
	"github.com/google/uuid"
)

type UserRequest struct {
	Email         string    `json:"email" binding:"required,email"`
	Password      string    `json:"password" binding:"required,min=8,containsany=!@$#*"`
	Name          string    `json:"name" binding:"max=128"`
	Role          string    `json:"role" binding:"required,oneof=manager professor student"`
	Admin         bool      `json:"admin"`
	InstitutionID uuid.UUID `json:"institution_id" binding:"required"`

	Department string `json:"department" binding:"required_if=Role professor"`

	Course      string    `json:"course" binding:"required_if=Role student"`
	TypeStudent string    `json:"type_student" binding:"required_if=Role student"`
	ProfessorID uuid.UUID `json:"professor_id" binding:"required_if=Role student"`
}
