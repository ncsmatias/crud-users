package request

import "github.com/google/uuid"

type InstitutionRequest struct {
	InstitutionType string    `json:"institution_type" binding:"required"`
	Name            string    `json:"name" binding:"required"`
	Phone           string    `json:"phone" binding:"max=128"`
	AddressID       uuid.UUID `json:"address_id" binding:"required"`
}
