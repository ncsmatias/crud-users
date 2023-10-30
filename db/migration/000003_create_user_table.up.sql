create table "users" (
  "user_id" uuid primary key default uuid_generate_v4(),
  "email" varchar unique not null,
  "name" varchar,
  "password" varchar not null,
  "role" varchar not null,
  "admin" boolean not null default false,
  "is_active" boolean not null default true,
  "institution_id" uuid REFERENCES institutions ("institution_id"),
  "created_at" timestamptz not null default (now())
);