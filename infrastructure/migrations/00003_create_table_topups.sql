-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "topups" (
  "top_up_id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  "user_id" uuid NOT NULL,
  "top_up_amount" numeric NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);


ALTER TABLE "topups" ADD CONSTRAINT "fk_user_id_topups" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "topups" DROP CONSTRAINT "fk_user_id_topups";

DROP TABLE IF EXISTS "topups";

-- +goose StatementEnd
