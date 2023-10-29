package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type StudentDomainInterface interface {
	GetCourse() string
	GetStudentType() string
	GetProfessorID() uuid.UUID
	ToString() string
}

func NewStudentDomain(course, studentType string, professorID uuid.UUID) StudentDomainInterface {
	return &studentDomain{course: course, studentType: studentType, professorID: professorID}
}

type studentDomain struct {
	course      string
	studentType string
	professorID uuid.UUID
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

func (sd *studentDomain) ToString() string {
	return fmt.Sprintf("{Course: %s, StudentType: %s, ProfessorID: %s}", sd.course, sd.studentType, sd.professorID)
}
