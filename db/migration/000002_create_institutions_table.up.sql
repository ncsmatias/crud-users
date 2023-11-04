create table "institutions"(
  "institution_id" uuid primary key default uuid_generate_v4(),
  "institution_type" varchar not null,
  "name" varchar not null,
  "phone" varchar,
  "zip_code" varchar,
  "street" varchar,
  "number" varchar,
  "city" varchar,
  "state" varchar,
  "UF" varchar,
  "country" varchar,
  "country_code" varchar,
  "address_id" uuid REFERENCES addresses ("address_id"),
  "created_at" timestamptz not null default (now())
);