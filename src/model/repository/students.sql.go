// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: students.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createStudent = `-- name: CreateStudent :one
insert into "students" (
  course,
  student_type,
  professor_id,
  user_id
) values ($1, $2, $3, $4) returning student_id
`

type CreateStudentParams struct {
	Course      sql.NullString `json:"course"`
	StudentType sql.NullString `json:"student_type"`
	ProfessorID uuid.NullUUID  `json:"professor_id"`
	UserID      uuid.NullUUID  `json:"user_id"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createStudent,
		arg.Course,
		arg.StudentType,
		arg.ProfessorID,
		arg.UserID,
	)
	var student_id uuid.UUID
	err := row.Scan(&student_id)
	return student_id, err
}

const findStudentByID = `-- name: FindStudentByID :one
select (
  course,
  student_type,
  professor_id,
  user_id
) from "students"
where student_id = $1 limit 1
`

func (q *Queries) FindStudentByID(ctx context.Context, studentID uuid.UUID) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, findStudentByID, studentID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const findStudentsByProfessorID = `-- name: FindStudentsByProfessorID :many
select (
  course,
  student_type,
  professor_id,
  user_id
) from "students"
where professor_id = $1
`

func (q *Queries) FindStudentsByProfessorID(ctx context.Context, professorID uuid.NullUUID) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, findStudentsByProfessorID, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []interface{}{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStudent = `-- name: UpdateStudent :one
update "students"
set course = $1, student_type = $2, professor_id = $3
where student_id = $4 returning student_id, course, student_type, user_id, professor_id, created_at
`

type UpdateStudentParams struct {
	Course      sql.NullString `json:"course"`
	StudentType sql.NullString `json:"student_type"`
	ProfessorID uuid.NullUUID  `json:"professor_id"`
	StudentID   uuid.UUID      `json:"student_id"`
}

func (q *Queries) UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateStudent,
		arg.Course,
		arg.StudentType,
		arg.ProfessorID,
		arg.StudentID,
	)
	var i Student
	err := row.Scan(
		&i.StudentID,
		&i.Course,
		&i.StudentType,
		&i.UserID,
		&i.ProfessorID,
		&i.CreatedAt,
	)
	return i, err
}
