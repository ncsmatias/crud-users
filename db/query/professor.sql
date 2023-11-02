-- name: CreateProfessor :one
insert into "professors" (
  department,
  user_id
) values ($1, $2) returning professor_id;

-- name: UpdateProfessor :one
update "professors"
set department = $1 
where professor_id = $2 returning *;

-- name: FindProfessorByID :one
select (professor_id, department, user_id) 
from "professors"
where professor_id = $1 limit 1;

-- name: FindProfessorByUserID :one
select (professor_id, department, user_id) 
from "professors"
where user_id = $1 limit 1;
