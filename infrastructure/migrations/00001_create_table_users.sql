-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "type_transaction" AS ENUM (
  'CREDIT',
  'DEBIT'
);

CREATE TYPE "type_source" AS ENUM (
  'TOPUP',
  'PAYMENT',
  'TRANSFER'
);

CREATE TABLE IF NOT EXISTS "users" (
  "user_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "phone_number" varchar UNIQUE NOT NULL,
  "address" text NOT NULL,
  "pin" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TYPE "type_transaction";

DROP TYPE "type_source";

DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
