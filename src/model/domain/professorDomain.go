package domain

import "fmt"

type ProfessorDomainInterface interface {
	GetDepartment() string
	GetUserID() string

	ToString() string
}

func NewProfessorDomain(department, userID string) ProfessorDomainInterface {
	return &professorDomain{department: department, userID: userID}
}

type professorDomain struct {
	department string
	userID     string
}

func (pd *professorDomain) GetDepartment() string {
	return pd.department
}

func (pd *professorDomain) GetUserID() string {
	return pd.userID
}

func (pd *professorDomain) ToString() string {
	return fmt.Sprintf("{Department: %s, UserID: %s}", pd.department, pd.userID)
}
