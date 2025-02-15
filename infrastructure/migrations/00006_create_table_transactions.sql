-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "transactions" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "user_id" uuid NOT NULL,
  "status" varchar NOT NULL,
  "transaction_type" type_transaction NOT NULL,
  "source_id" uuid,
  "source_type" type_source NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "transactions" ADD CONSTRAINT "fk_user_id_transactions" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "transactions" DROP CONSTRAINT "fk_user_id_transactions";


DROP TABLE IF EXISTS "transactions";
-- +goose StatementEnd
