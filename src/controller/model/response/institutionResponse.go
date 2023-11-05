package response

import "github.com/google/uuid"

type InstitutionResponse struct {
	ID              uuid.UUID
	InstitutionType string
	Name            string
	Phone           string
	ZipCode         string
	Street          string
	Number          string
	Neighborhood    string
	City            string
	State           string
	UF              string
	Country         string
	CountryCode     string
}
