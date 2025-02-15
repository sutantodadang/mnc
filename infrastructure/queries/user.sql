-- name: InsertUser :one
INSERT INTO users(first_name, last_name, phone_number, address, pin) VALUES($1, $2, $3, $4, $5) RETURNING *;

-- name: SelectOneUserByPhoneNumber :one
SELECT user_id, first_name, last_name, phone_number, address, pin, created_at FROM users WHERE phone_number = $1;

-- name: SelectOneUserById :one
SELECT a.user_id, a.first_name, a.last_name, a.phone_number, a.address, b.balance_id, b.balance_amount
FROM users a
JOIN balances b ON b.user_id = a.user_id
WHERE a.user_id = $1;

-- name: UpdateUser :one
UPDATE users SET first_name = $1, last_name = $2, address = $3, updated_at = now() WHERE user_id = $4 RETURNING *;