-- name: InsertPayment :one
INSERT INTO payments(user_id, remarks, payment_amount) VALUES($1, $2, $3) RETURNING *;
