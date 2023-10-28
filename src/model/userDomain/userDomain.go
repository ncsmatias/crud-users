package userdomain

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/ncsmatias/crud-users/src/configuration/resterr"
)

func NewUserDomain(email, name, password string, admin bool) UserDomainInterface {
	return &userDomain{Email: email, Name: name, Password: password, Admin: admin}
}

type userDomain struct {
	Email    string
	Password string
	Name     string
	Admin    bool
}

type UserDomainInterface interface {
	CreateUser() *resterr.RestErr
	UpdateUser(string) *resterr.RestErr
	FindUser(string) (*userDomain, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
