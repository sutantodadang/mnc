-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "balances_histories" (
  "balance_history_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "balance_id" uuid NOT NULL,
  "transaction_id" uuid,
  "balance_amount_before" numeric NOT NULL DEFAULT 0,
  "balance_amount_after" numeric NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "balances_histories" ADD CONSTRAINT "fk_balance_id_balances" FOREIGN KEY ("balance_id") REFERENCES "balances" ("balance_id");

ALTER TABLE "balances_histories" ADD CONSTRAINT "fk_transaction_id_transactions" FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("transaction_id");


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "balances_histories" DROP CONSTRAINT "fk_balance_id_balances";

ALTER TABLE "balances_histories" DROP CONSTRAINT "fk_transaction_id_transactions";

DROP TABLE IF EXISTS "balances_histories";

-- +goose StatementEnd
