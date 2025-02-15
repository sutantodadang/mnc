-- name: InsertBalanceHistory :exec
INSERT INTO balances_histories(balance_id, transaction_id, balance_amount_before, balance_amount_after) VALUES($1, $2, $3, $4);
