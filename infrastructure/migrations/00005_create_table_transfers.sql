-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "transfers" (
  "transfer_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "source_user_id" uuid NOT NULL,
  "target_user_id" uuid NOT NULL,
  "remarks" varchar NOT NULL,
  "transfer_amount" numeric NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "transfers" ADD CONSTRAINT "fk_source_user_id_transfers" FOREIGN KEY ("source_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "transfers" ADD CONSTRAINT "fk_target_user_id_transfers" FOREIGN KEY ("target_user_id") REFERENCES "users" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "transfers" DROP CONSTRAINT "fk_source_user_id_transfers";

ALTER TABLE "transfers" DROP CONSTRAINT "fk_target_user_id_transfers";

DROP TABLE IF EXISTS "transfers";

-- +goose StatementEnd
