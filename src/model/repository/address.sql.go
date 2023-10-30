// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: address.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAddress = `-- name: CreateAddress :one
insert into "addresses" (
  street,
  neighborhood,
  city,
  state,
  zip_code,
  notes
) values (
  $1, $2, $3, $4, $5, $6
) returning address_id
`

type CreateAddressParams struct {
	Street       string         `json:"street"`
	Neighborhood string         `json:"neighborhood"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	ZipCode      string         `json:"zip_code"`
	Notes        sql.NullString `json:"notes"`
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createAddress,
		arg.Street,
		arg.Neighborhood,
		arg.City,
		arg.State,
		arg.ZipCode,
		arg.Notes,
	)
	var address_id uuid.UUID
	err := row.Scan(&address_id)
	return address_id, err
}

const getAddressByID = `-- name: GetAddressByID :one
select (
  address_id,
  street,
  neighborhood,
  city,
  state,
  zip_code,
  notes
) from "addresses" where address_id = $1 limit 1
`

func (q *Queries) GetAddressByID(ctx context.Context, addressID uuid.UUID) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getAddressByID, addressID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getAddressByZIPCode = `-- name: GetAddressByZIPCode :one
select (
  address_id,
  street,
  neighborhood,
  city,
  state,
  zip_code,
  notes
) from "addresses" where zip_code = $1 limit 1
`

func (q *Queries) GetAddressByZIPCode(ctx context.Context, zipCode string) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getAddressByZIPCode, zipCode)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const updateAddress = `-- name: UpdateAddress :one
update "addresses"
set street = $1, neighborhood = $2, city = $3, state = $4, zip_code = $5, notes = $6
where address_id = $7 returning address_id, street, neighborhood, city, state, zip_code, notes, created_at
`

type UpdateAddressParams struct {
	Street       string         `json:"street"`
	Neighborhood string         `json:"neighborhood"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	ZipCode      string         `json:"zip_code"`
	Notes        sql.NullString `json:"notes"`
	AddressID    uuid.UUID      `json:"address_id"`
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, updateAddress,
		arg.Street,
		arg.Neighborhood,
		arg.City,
		arg.State,
		arg.ZipCode,
		arg.Notes,
		arg.AddressID,
	)
	var i Address
	err := row.Scan(
		&i.AddressID,
		&i.Street,
		&i.Neighborhood,
		&i.City,
		&i.State,
		&i.ZipCode,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}
