package domain

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	globaltypes "github.com/ncsmatias/crud-users/src/configuration/globalTypes"
)

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetRole() globaltypes.UserRole
	GetAdmin() bool
	GetinstitutionID() uuid.UUID

	SetID(uuid.UUID)

	GetJSONValue() (string, error)
	EncryptPassword()
	ToString() string
}

func NewUserDomain(email, name, password string, role globaltypes.UserRole, admin bool, institutionID uuid.UUID) UserDomainInterface {
	return &userDomain{Email: email, Name: name, Password: password, Role: role, Admin: admin, InstitutionID: institutionID}
}

type userDomain struct {
	ID            uuid.UUID
	Email         string
	Name          string
	Password      string
	Role          globaltypes.UserRole
	Admin         bool
	InstitutionID uuid.UUID
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetRole() globaltypes.UserRole {
	return ud.Role
}

func (ud *userDomain) GetAdmin() bool {
	return ud.Admin
}

func (ud *userDomain) GetinstitutionID() uuid.UUID {
	return ud.InstitutionID
}

func (ud *userDomain) SetID(userID uuid.UUID) {
	ud.ID = userID
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (ud *userDomain) ToString() string {
	adminStr := "false"
	if ud.Admin {
		adminStr = "true"
	}

	return fmt.Sprintf("{Email: %s, Name: %s, Role: %s, Admin: %s, institutionID: %s}", ud.Email, ud.Name, ud.Role, adminStr, ud.InstitutionID)
}
