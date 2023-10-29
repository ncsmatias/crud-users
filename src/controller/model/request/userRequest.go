package request

import (
	"github.com/google/uuid"
	globaltypes "github.com/ncsmatias/crud-users/src/configuration/globalTypes"
)

type UserRequest struct {
	Email         string               `json:"email" binding:"required,email"`
	Password      string               `json:"password" binding:"required,min=8,containsany=!@$#*"`
	Name          string               `json:"name" binding:"max=128"`
	Role          globaltypes.UserRole `json:"role" binding:"required,oneof=manager professor student"`
	Admin         bool                 `json:"admin" binding:"required"`
	InstitutionID uuid.UUID            `json:"institution_id" binding:"required"`

	Department string `json:"department" binding:"required_if=Role professor"`

	Course      string    `json:"course" binding:"required_if=Role student"`
	TypeStudent string    `json:"type_student" binding:"required_if=Role student,oneof=undergraduate masters doctoral"`
	ProfessorID uuid.UUID `json:"professor_id" binding:"required_if=Role student"`
}
