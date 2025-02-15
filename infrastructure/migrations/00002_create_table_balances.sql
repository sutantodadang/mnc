-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "balances" (
  "balance_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "user_id" uuid NOT NULL,
  "balance_amount" numeric NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "balances" ADD CONSTRAINT "fk_user_id_balances" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "balances" DROP CONSTRAINT "fk_user_id_balances";


DROP TABLE IF EXISTS "balances";


-- +goose StatementEnd
