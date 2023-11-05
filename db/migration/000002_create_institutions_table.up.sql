create table "institutions"(
  "institution_id" uuid primary key default uuid_generate_v4(),
  "institution_type" varchar not null,
  "name" varchar not null,
  "phone" varchar,
  "zip_code" varchar,
  "street" varchar,
  "number" varchar,
  "neighborhood" varchar,
  "city" varchar,
  "state" varchar,
  "uf" varchar,
  "country" varchar,
  "country_code" varchar,
  "is_active" boolean not null default true,
  "created_at" timestamptz not null default (now())
);