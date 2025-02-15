-- name: InsertTopUp :one
INSERT INTO topups(user_id, top_up_amount) VALUES($1, $2) RETURNING *;

