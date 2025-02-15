-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS"payments" (
  "payment_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "user_id" uuid NOT NULL,
  "remarks" varchar NOT NULL,
  "payment_amount" numeric NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);


ALTER TABLE "payments" ADD CONSTRAINT "fk_user_id_payments" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "payments" DROP CONSTRAINT "fk_user_id_payments";

DROP TABLE IF EXISTS "payments";

-- +goose StatementEnd
