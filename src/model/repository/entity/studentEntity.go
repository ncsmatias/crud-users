package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type StudentEntity struct {
	ID          uuid.UUID
	Course      sql.NullString
	StudentType sql.NullString
	ProfessorID uuid.NullUUID
	UserID      uuid.NullUUID
}
