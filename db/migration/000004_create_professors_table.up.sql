create table "professors" (
  "professor_id" uuid primary key default uuid_generate_v4(),
  "department" varchar not null,
  "user_id" uuid REFERENCES users ("user_id"),
  "created_at" timestamptz not null default (now())
)