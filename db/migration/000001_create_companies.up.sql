create table companies (
  "ID" bigserial PRIMARY KEY,
  "Name" varchar NOT NULL,
  "Zip" varchar(5) NOT NULL,
  "Website" varchar,
  "CreateAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdateAt" timestamptz NOT NULL DEFAULT (now())
)