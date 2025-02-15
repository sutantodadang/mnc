-- name: InsertDefaultBalance :exec
INSERT INTO balances(user_id) VALUES($1);

-- name: UpdateBalance :exec
UPDATE balances SET balance_amount = $2 WHERE balance_id = $1;