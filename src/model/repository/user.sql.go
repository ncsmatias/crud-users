// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
insert into "users" (
  email, 
  name, 
  password, 
  role, 
  admin, 
  institution_id
) values (
  $1, $2, $3, $4, $5, $6
) returning user_id
`

type CreateUserParams struct {
	Email         string         `json:"email"`
	Name          sql.NullString `json:"name"`
	Password      string         `json:"password"`
	Role          string         `json:"role"`
	Admin         bool           `json:"admin"`
	InstitutionID uuid.NullUUID  `json:"institution_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Name,
		arg.Password,
		arg.Role,
		arg.Admin,
		arg.InstitutionID,
	)
	var user_id uuid.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select (
  email, 
  name, 
  password, 
  role, 
  admin, 
  is_active, 
  institution_id
) from "users" where email = $1 limit 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getUserByID = `-- name: GetUserByID :one
select (
  email, 
  name, 
  password, 
  role, 
  admin, 
  is_active, 
  institution_id
) from "users" where user_id = $1 limit 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID uuid.UUID) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const updateUser = `-- name: UpdateUser :one
update "users"
set email = $1, name = $2, password = $3, role = $4, admin = $5, is_active = $6, institution_id = $7
where user_id = $8 returning user_id, email, name, password, role, admin, is_active, institution_id, created_at
`

type UpdateUserParams struct {
	Email         string         `json:"email"`
	Name          sql.NullString `json:"name"`
	Password      string         `json:"password"`
	Role          string         `json:"role"`
	Admin         bool           `json:"admin"`
	IsActive      bool           `json:"is_active"`
	InstitutionID uuid.NullUUID  `json:"institution_id"`
	UserID        uuid.UUID      `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Email,
		arg.Name,
		arg.Password,
		arg.Role,
		arg.Admin,
		arg.IsActive,
		arg.InstitutionID,
		arg.UserID,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Name,
		&i.Password,
		&i.Role,
		&i.Admin,
		&i.IsActive,
		&i.InstitutionID,
		&i.CreatedAt,
	)
	return i, err
}
