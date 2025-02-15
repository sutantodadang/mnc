-- name: InsertTransaction :one
INSERT INTO transactions(user_id, status, transaction_type, source_id, source_type ) VALUES($1, $2, $3, $4, $5) RETURNING *;

-- name: SelectTransactionByUserId :many
SELECT t.transaction_id, t.user_id, t.status, t.transaction_type, t.source_id, t.source_type, t.created_at,
bh.balance_amount_before, bh.balance_amount_after,
t2.top_up_id, t2.top_up_amount,
p.payment_id, p.payment_amount, p.remarks AS payment_remarks,
t3.transfer_id, t3.transfer_amount, t3.remarks AS transfer_remarks
FROM transactions t 
LEFT JOIN balances_histories bh ON bh.transaction_id = t.transaction_id 
LEFT JOIN topups t2 ON t.source_id = t2.top_up_id 
LEFT JOIN payments p ON t.source_id = p.payment_id 
LEFT JOIN transfers t3 ON t.source_id = t3.transfer_id 
WHERE t.user_id = $1;