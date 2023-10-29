create table "institutions"(
  "ID" uuid primary key default uuid_generate_v4(),
  "name" varchar not null,
  "phone" varchar
);