create table "addresses" (
  "address_id" uuid primary key default uuid_generate_v4(),
  "street" varchar not null,
  "neighborhood" varchar not null,
  "city" varchar not null,
  "state" varchar not null,
  "zip_code" varchar not null,
  "notes" varchar,
  "created_at" timestamptz not null default (now())
)