-- name: CreateInstitution :one
insert into "institutions" (
  institution_type,
  name,
  phone,
  zip_code,
  street,
  number,
  neighborhood,
  city,
  state,
  UF,
  country,
  country_code
) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
returning institution_id;

-- name: UpdateInstitution :one
update "institutions"
set institution_type = $1, 
    name = $2, 
    phone = $3,
    zip_code = $4,
    street = $5,
    number = $6,
    city = $7,
    neighborhood = $8,
    state = $9,
    UF = $10,
    country = $11,
    country_code = $12
where institution_id = $13 returning *;

-- name: GetInstitutionByName :many
select (
  institution_id,
  institution_type,
  name,
  phone,
  zip_code,
  street,
  number,
  city,
  neighborhood,
  state,
  UF,
  country,
  country_code
) from "institutions"
where name like $1;

-- name: GetInstitutionByID :one
select (
  institution_id,
  institution_type,
  name,
  phone,
  zip_code,
  street,
  number,
  city,
  neighborhood,
  state,
  UF,
  country,
  country_code
) from "institutions"
where institution_id = $1 limit 1;