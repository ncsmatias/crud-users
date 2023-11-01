package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type AddressEntity struct {
	ID           uuid.UUID
	Street       string
	Neighborhood string
	City         string
	State        string
	ZipCode      string
	Notes        sql.NullString
}
