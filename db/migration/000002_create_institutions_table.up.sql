create table "institutions"(
  "institution_id" uuid primary key default uuid_generate_v4(),
  "name" varchar not null,
  "phone" varchar,
  "address_id" uuid REFERENCES addresses ("address_id"),
  "created_at" timestamptz not null default (now())
);