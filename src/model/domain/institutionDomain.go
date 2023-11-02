package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type InstitutionDomainInterface interface {
	GetID() uuid.UUID
	GetInstitutionType() string
	GetName() string
	GetPhone() string
	GetAddressID() uuid.UUID

	SetID(uuid.UUID)

	ToString() string
}

func NewInstitutionDomain(institutionType, name, phone string, addressID uuid.UUID) InstitutionDomainInterface {
	return &institutionDomain{
		institutionType: institutionType,
		name:            name,
		phone:           phone,
		addressID:       addressID,
	}
}

type institutionDomain struct {
	id              uuid.UUID
	institutionType string
	name            string
	phone           string
	addressID       uuid.UUID
}

func (instd *institutionDomain) SetID(institutionID uuid.UUID) {
	instd.id = institutionID
}

func (instd *institutionDomain) GetAddressID() uuid.UUID {
	return instd.addressID
}

func (instd *institutionDomain) GetID() uuid.UUID {
	return instd.id
}

func (instd *institutionDomain) GetInstitutionType() string {
	return instd.institutionType
}

func (instd *institutionDomain) GetName() string {
	return instd.name
}

func (instd *institutionDomain) GetPhone() string {
	return instd.phone
}

func (instd *institutionDomain) ToString() string {
	return fmt.Sprintf("ID: %s, InstitutionType: %s, Name: %s, Phone: %s, AddressID: %s", instd.id,
		instd.institutionType,
		instd.name,
		instd.phone,
		instd.addressID,
	)
}
