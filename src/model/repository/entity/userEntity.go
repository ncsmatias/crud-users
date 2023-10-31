package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID            uuid.UUID
	Email         string
	Name          sql.NullString
	Password      string
	Role          string
	Admin         bool
	InstitutionID uuid.NullUUID
}
