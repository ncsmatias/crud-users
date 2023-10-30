create table "students" (
  "student_id" uuid primary key default uuid_generate_v4(),
  "course" varchar,
  "student_type" varchar,
  "user_id" uuid REFERENCES users ("user_id"),
  "professor_id" uuid REFERENCES professors ("professor_id"),
  "created_at" timestamptz not null default (now())
)