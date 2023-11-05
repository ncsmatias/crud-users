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
	ZipCode         sql.NullString
	Street          sql.NullString
	Number          sql.NullString
	Neighborhood    sql.NullString
	City            sql.NullString
	State           sql.NullString
	UF              sql.NullString
	Country         sql.NullString
	CountryCode     sql.NullString
}
