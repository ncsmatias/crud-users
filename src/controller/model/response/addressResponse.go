package response

import "github.com/google/uuid"

type AddressResponse struct {
	ID           uuid.UUID
	Street       string
	Neighborhood string
	City         string
	State        string
	ZipCode      string
	Notes        string
}
