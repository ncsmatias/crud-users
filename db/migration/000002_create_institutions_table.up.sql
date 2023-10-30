create table "institutions"(
  "institution_id" uuid primary key default uuid_generate_v4(),
  "institution_type" varchar not null,
  "name" varchar not null,
  "phone" varchar,
  "address_id" uuid REFERENCES addresses ("address_id"),
  "created_at" timestamptz not null default (now())
);