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
	GetZipCode() string
	GetStreet() string
	GetNumber() string
	GetNeighborhood() string
	GetCity() string
	GetState() string
	GetUF() string
	GetCountry() string
	GetCountryCode() string

	SetID(uuid.UUID)

	ToString() string
}

func NewInstitutionDomain(
	institutionType,
	name,
	phone,
	zipCode,
	street,
	number,
	neighborhood,
	city,
	state,
	uf,
	country,
	countryCode string,
) InstitutionDomainInterface {
	return &institutionDomain{
		institutionType: institutionType,
		name:            name,
		phone:           phone,
		zipCode:         zipCode,
		street:          street,
		number:          number,
		neighborhood:    neighborhood,
		city:            city,
		state:           state,
		uf:              uf,
		country:         country,
		countryCode:     countryCode,
	}
}

type institutionDomain struct {
	id              uuid.UUID
	institutionType string
	name            string
	phone           string
	zipCode         string
	street          string
	number          string
	neighborhood    string
	city            string
	state           string
	uf              string
	country         string
	countryCode     string
}

func (instd *institutionDomain) GetCity() string {
	return instd.city
}

func (instd *institutionDomain) GetCountry() string {
	return instd.country
}

func (instd *institutionDomain) GetCountryCode() string {
	return instd.countryCode
}

func (instd *institutionDomain) GetNeighborhood() string {
	return instd.neighborhood
}

func (instd *institutionDomain) GetNumber() string {
	return instd.number
}

func (instd *institutionDomain) GetState() string {
	return instd.state
}

func (instd *institutionDomain) GetStreet() string {
	return instd.street
}

func (instd *institutionDomain) GetUF() string {
	return instd.uf
}

func (instd *institutionDomain) GetZipCode() string {
	return instd.zipCode
}

func (instd *institutionDomain) SetID(institutionID uuid.UUID) {
	instd.id = institutionID
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
	return fmt.Sprintf("ID: %s, InstitutionType: %s, Name: %s, Phone: %s", instd.id,
		instd.institutionType,
		instd.name,
		instd.phone,
	)
}
