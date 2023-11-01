package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type AddressDomainInterface interface {
	GetID() uuid.UUID
	GetStreet() string
	GetNeighborhood() string
	GetCity() string
	GetState() string
	GetZipCode() string
	GetNotes() string

	SetID(uuid.UUID)

	ToString() string
}

func NewAddressDomain(street, neighborhood, city, state, zipCode, notes string) AddressDomainInterface {
	return &addressDomain{
		street:       state,
		neighborhood: neighborhood,
		city:         city,
		state:        state,
		zipCode:      zipCode,
		notes:        notes,
	}
}

type addressDomain struct {
	id           uuid.UUID
	street       string
	neighborhood string
	city         string
	state        string
	zipCode      string
	notes        string
}

func (ad *addressDomain) GetCity() string {
	return ad.city
}

func (ad *addressDomain) GetID() uuid.UUID {
	return ad.id
}

func (ad *addressDomain) GetNeighborhood() string {
	return ad.neighborhood
}

func (ad *addressDomain) GetNotes() string {
	return ad.notes
}

func (ad *addressDomain) GetState() string {
	return ad.state
}

func (ad *addressDomain) GetStreet() string {
	return ad.street
}

func (ad *addressDomain) GetZipCode() string {
	return ad.zipCode
}

func (ad *addressDomain) SetID(addressID uuid.UUID) {
	ad.id = addressID
}

func (ad *addressDomain) ToString() string {

	return fmt.Sprintf("ID: %s, street: %s, neighborhood: %s, city: %s, state: %s, zipCode: %s, notes: %s", ad.id,
		ad.street,
		ad.neighborhood,
		ad.city,
		ad.state,
		ad.zipCode,
		ad.notes,
	)
}
