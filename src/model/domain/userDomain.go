package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAdmin() bool

	EncryptPassword()
	ToString() string
}

func NewUserDomain(email, name, password string, admin bool) UserDomainInterface {
	return &userDomain{email: email, name: name, password: password, admin: admin}
}

type userDomain struct {
	email    string
	name     string
	password string
	admin    bool
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

func (ud *userDomain) GetAdmin() bool {
	return ud.admin
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

	return fmt.Sprintf("{Email: %s, Name: %s, Admin: %s}", ud.email, ud.name, adminStr)
}
