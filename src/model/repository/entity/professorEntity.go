package entity

import (
	"github.com/google/uuid"
)

type ProfessorEntity struct {
	ID         uuid.UUID
	Department string
	UserID     uuid.NullUUID
}
