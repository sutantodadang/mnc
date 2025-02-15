-- name: InsertTransfer :one
INSERT INTO transfers(source_user_id, target_user_id, remarks, transfer_amount) VALUES($1, $2, $3, $4) RETURNING *;
