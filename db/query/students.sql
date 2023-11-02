-- name: CreateStudent :one
insert into "students" (
  course,
  student_type,
  professor_id,
  user_id
) values ($1, $2, $3, $4) returning student_id;

-- name: UpdateStudent :one
update "students"
set course = $1, student_type = $2, professor_id = $3
where student_id = $4 returning *;

-- name: FindStudentByID :one
select (
  course,
  student_type,
  professor_id,
  user_id
) from "students"
where student_id = $1 limit 1;

-- name: FindStudentsByProfessorID :many
select (
  course,
  student_type,
  professor_id,
  user_id
) from "students"
where professor_id = $1;