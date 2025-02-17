// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transfer.sql

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertTransfer = `-- name: InsertTransfer :one
INSERT INTO transfers(source_user_id, target_user_id, remarks, transfer_amount) VALUES($1, $2, $3, $4) RETURNING transfer_id, source_user_id, target_user_id, remarks, transfer_amount, created_at, updated_at
`

type InsertTransferParams struct {
	SourceUserID   pgtype.UUID    `db:"source_user_id" json:"source_user_id"`
	TargetUserID   pgtype.UUID    `db:"target_user_id" json:"target_user_id"`
	Remarks        string         `db:"remarks" json:"remarks"`
	TransferAmount pgtype.Numeric `db:"transfer_amount" json:"transfer_amount"`
}

func (q *Queries) InsertTransfer(ctx context.Context, arg InsertTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, insertTransfer,
		arg.SourceUserID,
		arg.TargetUserID,
		arg.Remarks,
		arg.TransferAmount,
	)
	var i Transfer
	err := row.Scan(
		&i.TransferID,
		&i.SourceUserID,
		&i.TargetUserID,
		&i.Remarks,
		&i.TransferAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
