package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type InstitutionEntity struct {
	ID              uuid.UUID
	InstitutionType string
	Name            string
	Phone           sql.NullString
	AddressID       uuid.NullUUID
}
