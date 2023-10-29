package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ProfessorDomainInterface interface {
	GetDepartment() string
	GetUserID() uuid.UUID

	ToString() string
}

func NewProfessorDomain(department string, userID uuid.UUID) ProfessorDomainInterface {
	return &professorDomain{department: department, userID: userID}
}

type professorDomain struct {
	department string
	userID     uuid.UUID
}

func (pd *professorDomain) GetDepartment() string {
	return pd.department
}

func (pd *professorDomain) GetUserID() uuid.UUID {
	return pd.userID
}

func (pd *professorDomain) ToString() string {
	return fmt.Sprintf("{Department: %s, UserID: %s}", pd.department, pd.userID)
}
