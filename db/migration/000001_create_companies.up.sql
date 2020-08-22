create table companies (
  "id" bigserial PRIMARY KEY,
  "company_name" varchar NOT NULL,
  "zipcode" varchar(5) NOT NULL,
  "website" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT NULL
)
