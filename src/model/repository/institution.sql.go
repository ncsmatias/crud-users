// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: institution.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createInstitution = `-- name: CreateInstitution :one
insert into "institutions" (
  institution_type,
  name,
  phone,
  address_id
) values ($1, $2, $3, $4) returning institution_id
`

type CreateInstitutionParams struct {
	InstitutionType string         `json:"institution_type"`
	Name            string         `json:"name"`
	Phone           sql.NullString `json:"phone"`
	AddressID       uuid.NullUUID  `json:"address_id"`
}

func (q *Queries) CreateInstitution(ctx context.Context, arg CreateInstitutionParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createInstitution,
		arg.InstitutionType,
		arg.Name,
		arg.Phone,
		arg.AddressID,
	)
	var institution_id uuid.UUID
	err := row.Scan(&institution_id)
	return institution_id, err
}

const getInstitutionByID = `-- name: GetInstitutionByID :one
select (
  institution_id,
  institution_type,
  name,
  phone,
  address_id
) from "institutions"
where institution_id = $1 limit 1
`

func (q *Queries) GetInstitutionByID(ctx context.Context, institutionID uuid.UUID) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getInstitutionByID, institutionID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getInstitutionByName = `-- name: GetInstitutionByName :many
select (
  institution_id,
  institution_type,
  name,
  phone,
  address_id
) from "institutions"
where name like $1 limit 1
`

func (q *Queries) GetInstitutionByName(ctx context.Context, name string) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, getInstitutionByName, name)
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

const updateInstitution = `-- name: UpdateInstitution :one
update "institutions"
set institution_type = $1, name = $2, phone = $3
where institution_id = $4 returning institution_id, institution_type, name, phone, address_id, created_at
`

type UpdateInstitutionParams struct {
	InstitutionType string         `json:"institution_type"`
	Name            string         `json:"name"`
	Phone           sql.NullString `json:"phone"`
	InstitutionID   uuid.UUID      `json:"institution_id"`
}

func (q *Queries) UpdateInstitution(ctx context.Context, arg UpdateInstitutionParams) (Institution, error) {
	row := q.db.QueryRowContext(ctx, updateInstitution,
		arg.InstitutionType,
		arg.Name,
		arg.Phone,
		arg.InstitutionID,
	)
	var i Institution
	err := row.Scan(
		&i.InstitutionID,
		&i.InstitutionType,
		&i.Name,
		&i.Phone,
		&i.AddressID,
		&i.CreatedAt,
	)
	return i, err
}