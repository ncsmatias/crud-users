-- name: CreateUser :one
insert into "users" (
  email, 
  name, 
  password, 
  role, 
  admin, 
  institution_id
) values (
  $1, $2, $3, $4, $5, $6
) returning user_id;

-- name: UpdateUser :one
update "users"
set email = $1, name = $2, password = $3, role = $4, admin = $5, is_active = $6, institution_id = $7
where user_id = $8 returning *;

-- name: GetUserByID :one
select (
  email, 
  name, 
  password, 
  role, 
  admin, 
  is_active, 
  institution_id
) from "users" where user_id = $1 limit 1;

-- name: GetUserByEmail :one
select (
  email, 
  name, 
  password, 
  role, 
  admin, 
  is_active, 
  institution_id
) from "users" where email = $1 limit 1;