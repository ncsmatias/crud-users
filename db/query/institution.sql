-- name: CreateInstitution :one
insert into "institutions" (
  institution_type,
  name,
  phone,
  address_id
) values ($1, $2, $3, $4) returning institution_id;

-- name: UpdateInstitution :one
update "institutions"
set institution_type = $1, name = $2, phone = $3
where institution_id = $4 returning *;

-- name: GetInstitutionByName :many
select (
  institution_id,
  institution_type,
  name,
  phone,
  address_id
) from "institutions"
where name like $1 limit 1;

-- name: GetInstitutionByID :one
select (
  institution_id,
  institution_type,
  name,
  phone,
  address_id
) from "institutions"
where institution_id = $1 limit 1;