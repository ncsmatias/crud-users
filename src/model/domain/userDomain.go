package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

type UserDomainInterface interface {
	GetID() uuid.UUID
	GetEmail() string
	GetName() string
	GetPassword() string
	GetRole() string
	GetAdmin() bool
	GetinstitutionID() uuid.UUID

	SetID(uuid.UUID)

	EncryptPassword()
	ToString() string
}

func NewUserDomain(email, name, password string, role string, admin bool, institutionID uuid.UUID) UserDomainInterface {
	return &userDomain{email: email, name: name, password: password, role: role, admin: admin, institutionID: institutionID}
}

type userDomain struct {
	id            uuid.UUID
	email         string
	name          string
	password      string
	role          string
	admin         bool
	institutionID uuid.UUID
}

func (ud *userDomain) GetID() uuid.UUID {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetRole() string {
	return ud.role
}

func (ud *userDomain) GetAdmin() bool {
	return ud.admin
}

func (ud *userDomain) GetinstitutionID() uuid.UUID {
	return ud.institutionID
}

func (ud *userDomain) SetID(userID uuid.UUID) {
	ud.id = userID
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) ToString() string {
	adminStr := "false"
	if ud.admin {
		adminStr = "true"
	}

	return fmt.Sprintf("{Email: %s, Name: %s, Role: %s, Admin: %s, institutionID: %s}", ud.email, ud.name, ud.role, adminStr, ud.institutionID)
}
