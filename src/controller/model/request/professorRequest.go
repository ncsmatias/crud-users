package request

import "github.com/google/uuid"

type ProfessorRequest struct {
	Department string    `json:"department" binding:"required"`
	UserID     uuid.UUID `json:"user_id" binding:"required"`
}
