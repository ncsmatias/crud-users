-- name: CreateAddress :one
insert into "addresses" (
  street,
  neighborhood,
  city,
  state,
  zip_code,
  notes
) values (
  $1, $2, $3, $4, $5, $6
) returning address_id;

-- name: UpdateAddress :one
update "addresses"
set street = $1, neighborhood = $2, city = $3, state = $4, zip_code = $5, notes = $6
where address_id = $7 returning *;

-- name: GetAddressByZIPCode :one
select (
  address_id,
  street,
  neighborhood,
  city,
  state,
  zip_code,
  notes
) from "addresses" where zip_code = $1 limit 1;

-- name: GetAddressByID :one
select (
  address_id,
  street,
  neighborhood,
  city,
  state,
  zip_code,
  notes
) from "addresses" where address_id = $1 limit 1;