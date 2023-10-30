-- name: CreateUser :one
insert into "users" (
  email, 
  name, 
  password, 
  role, 
  admin, 
  is_active, 
  institution_id
) values (
  $1, $2, $3, $4, $5, $6, $7
) returning user_id;

-- name: GetUserByID :one
select (
  email, 
  name, 
  password, 
  role, 
  admin, 
  is_active, 
  institution_id
) from "users" where user_id == $1 limit 1;