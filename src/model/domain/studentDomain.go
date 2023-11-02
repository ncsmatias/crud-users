package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type StudentDomainInterface interface {
	GetID() uuid.UUID
	GetCourse() string
	GetStudentType() string
	GetProfessorID() uuid.UUID
	GetUserID() uuid.UUID

	SetID(uuid.UUID)

	ToString() string
}

func NewStudentDomain(course, studentType string, professorID, userID uuid.UUID) StudentDomainInterface {
	return &studentDomain{course: course, studentType: studentType, professorID: professorID, userID: userID}
}

type studentDomain struct {
	id          uuid.UUID
	course      string
	studentType string
	professorID uuid.UUID
	userID      uuid.UUID
}

func (sd *studentDomain) SetID(studentID uuid.UUID) {
	sd.id = studentID
}

func (sd *studentDomain) GetID() uuid.UUID {
	return sd.id
}

func (sd *studentDomain) GetCourse() string {
	return sd.course
}

func (sd *studentDomain) GetStudentType() string {
	return sd.studentType
}

func (sd *studentDomain) GetProfessorID() uuid.UUID {
	return sd.professorID
}

func (sd *studentDomain) GetUserID() uuid.UUID {
	return sd.userID
}

func (sd *studentDomain) ToString() string {
	return fmt.Sprintf("{ID: %s, Course: %s, StudentType: %s, ProfessorID: %s}", sd.id, sd.course, sd.studentType, sd.professorID)
}
