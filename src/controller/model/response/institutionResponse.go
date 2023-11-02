package response

import "github.com/google/uuid"

type InstitutionRequest struct {
	ID              uuid.UUID
	InstitutionType string
	Name            string
	Phone           string
	AddressID       uuid.UUID
}
