// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: balance.sql

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertDefaultBalance = `-- name: InsertDefaultBalance :exec
INSERT INTO balances(user_id) VALUES($1)
`

func (q *Queries) InsertDefaultBalance(ctx context.Context, userID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, insertDefaultBalance, userID)
	return err
}

const updateBalance = `-- name: UpdateBalance :exec
UPDATE balances SET balance_amount = $2 WHERE balance_id = $1
`

type UpdateBalanceParams struct {
	BalanceID     pgtype.UUID    `db:"balance_id" json:"balance_id"`
	BalanceAmount pgtype.Numeric `db:"balance_amount" json:"balance_amount"`
}

func (q *Queries) UpdateBalance(ctx context.Context, arg UpdateBalanceParams) error {
	_, err := q.db.Exec(ctx, updateBalance, arg.BalanceID, arg.BalanceAmount)
	return err
}
