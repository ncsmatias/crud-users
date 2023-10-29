create table "professors" (
  "ID" uuid primary key default uuid_generate_v4(),
  "department" varchar not null,
  "user_id" uuid REFERENCES users ("ID"),
  "created_at" timestamptz not null default (now())
)