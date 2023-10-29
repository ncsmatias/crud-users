create table "students" (
  "ID" uuid primary key default uuid_generate_v4(),
  "course" varchar,
  "student_type" varchar,
  "user_id" uuid REFERENCES users ("ID"),
  "professors_id" uuid REFERENCES professors ("ID")
)