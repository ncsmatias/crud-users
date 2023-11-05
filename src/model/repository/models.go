// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Institution struct {
	InstitutionID   uuid.UUID      `json:"institution_id"`
	InstitutionType string         `json:"institution_type"`
	Name            string         `json:"name"`
	Phone           sql.NullString `json:"phone"`
	ZipCode         sql.NullString `json:"zip_code"`
	Street          sql.NullString `json:"street"`
	Number          sql.NullString `json:"number"`
	Neighborhood    sql.NullString `json:"neighborhood"`
	City            sql.NullString `json:"city"`
	State           sql.NullString `json:"state"`
	Uf              sql.NullString `json:"uf"`
	Country         sql.NullString `json:"country"`
	CountryCode     sql.NullString `json:"country_code"`
	IsActive        bool           `json:"is_active"`
	CreatedAt       time.Time      `json:"created_at"`
}

type Professor struct {
	ProfessorID uuid.UUID     `json:"professor_id"`
	Department  string        `json:"department"`
	UserID      uuid.NullUUID `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
}

type Student struct {
	StudentID   uuid.UUID      `json:"student_id"`
	Course      sql.NullString `json:"course"`
	StudentType sql.NullString `json:"student_type"`
	UserID      uuid.NullUUID  `json:"user_id"`
	ProfessorID uuid.NullUUID  `json:"professor_id"`
	CreatedAt   time.Time      `json:"created_at"`
}

type User struct {
	UserID        uuid.UUID      `json:"user_id"`
	Email         string         `json:"email"`
	Name          sql.NullString `json:"name"`
	Password      string         `json:"password"`
	Role          string         `json:"role"`
	Admin         bool           `json:"admin"`
	IsActive      bool           `json:"is_active"`
	InstitutionID uuid.NullUUID  `json:"institution_id"`
	CreatedAt     time.Time      `json:"created_at"`
}
