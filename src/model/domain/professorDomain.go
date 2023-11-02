package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ProfessorDomainInterface interface {
	GetID() uuid.UUID
	GetDepartment() string
	GetUserID() uuid.UUID

	SetID(uuid.UUID)
	ToString() string
}

func NewProfessorDomain(department string, userID uuid.UUID) ProfessorDomainInterface {
	return &professorDomain{department: department, userID: userID}
}

type professorDomain struct {
	id         uuid.UUID
	department string
	userID     uuid.UUID
}

func (pd *professorDomain) SetID(professorID uuid.UUID) {
	pd.id = professorID
}

func (pd *professorDomain) GetID() uuid.UUID {
	return pd.id
}

func (pd *professorDomain) GetDepartment() string {
	return pd.department
}

func (pd *professorDomain) GetUserID() uuid.UUID {
	return pd.userID
}

func (pd *professorDomain) ToString() string {
	return fmt.Sprintf("{ID: %s, Department: %s, UserID: %s}", pd.id, pd.department, pd.userID)
}
