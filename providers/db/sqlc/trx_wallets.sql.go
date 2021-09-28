// Code generated by sqlc. DO NOT EDIT.
// source: trx_wallets.sql

package db

import (
	"context"
	"database/sql"
)

const createTrxWallet = `-- name: CreateTrxWallet :one
INSERT INTO trx_wallets (
    wallet_id, 
    amount, 
    balance_before, 
    balance, 
    dest_model, 
    dest_model_id, 
    trx_receives_id, 
    debit,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING id, wallet_id, amount, balance_before, balance, dest_model, dest_model_id, trx_receives_id, created_at, created_by, updated_at, updated_by, debit
`

type CreateTrxWalletParams struct {
	WalletID      int32          `json:"wallet_id"`
	Amount        string         `json:"amount"`
	BalanceBefore string         `json:"balance_before"`
	Balance       string         `json:"balance"`
	DestModel     sql.NullString `json:"dest_model"`
	DestModelID   sql.NullInt32  `json:"dest_model_id"`
	TrxReceivesID int32          `json:"trx_receives_id"`
	Debit         bool           `json:"debit"`
	CreatedBy     int64          `json:"created_by"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) CreateTrxWallet(ctx context.Context, arg CreateTrxWalletParams) (TrxWallet, error) {
	row := q.db.QueryRowContext(ctx, createTrxWallet,
		arg.WalletID,
		arg.Amount,
		arg.BalanceBefore,
		arg.Balance,
		arg.DestModel,
		arg.DestModelID,
		arg.TrxReceivesID,
		arg.Debit,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i TrxWallet
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.BalanceBefore,
		&i.Balance,
		&i.DestModel,
		&i.DestModelID,
		&i.TrxReceivesID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Debit,
	)
	return i, err
}

const deleteTrxWallet = `-- name: DeleteTrxWallet :exec
DELETE FROM trx_wallets
WHERE id = $1
`

func (q *Queries) DeleteTrxWallet(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTrxWallet, id)
	return err
}

const getTrxWallet = `-- name: GetTrxWallet :one
SELECT id, wallet_id, amount, balance_before, balance, dest_model, dest_model_id, trx_receives_id, created_at, created_by, updated_at, updated_by, debit FROM trx_wallets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTrxWallet(ctx context.Context, id int32) (TrxWallet, error) {
	row := q.db.QueryRowContext(ctx, getTrxWallet, id)
	var i TrxWallet
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.BalanceBefore,
		&i.Balance,
		&i.DestModel,
		&i.DestModelID,
		&i.TrxReceivesID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Debit,
	)
	return i, err
}

const getTrxWalletForUpdate = `-- name: GetTrxWalletForUpdate :one
SELECT id, wallet_id, amount, balance_before, balance, dest_model, dest_model_id, trx_receives_id, created_at, created_by, updated_at, updated_by, debit FROM trx_wallets
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTrxWalletForUpdate(ctx context.Context, id int32) (TrxWallet, error) {
	row := q.db.QueryRowContext(ctx, getTrxWalletForUpdate, id)
	var i TrxWallet
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.BalanceBefore,
		&i.Balance,
		&i.DestModel,
		&i.DestModelID,
		&i.TrxReceivesID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Debit,
	)
	return i, err
}

const listTrxWallets = `-- name: ListTrxWallets :many
SELECT id, wallet_id, amount, balance_before, balance, dest_model, dest_model_id, trx_receives_id, created_at, created_by, updated_at, updated_by, debit FROM trx_wallets
WHERE wallet_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTrxWalletsParams struct {
	WalletID int32 `json:"wallet_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListTrxWallets(ctx context.Context, arg ListTrxWalletsParams) ([]TrxWallet, error) {
	rows, err := q.db.QueryContext(ctx, listTrxWallets, arg.WalletID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TrxWallet
	for rows.Next() {
		var i TrxWallet
		if err := rows.Scan(
			&i.ID,
			&i.WalletID,
			&i.Amount,
			&i.BalanceBefore,
			&i.Balance,
			&i.DestModel,
			&i.DestModelID,
			&i.TrxReceivesID,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.Debit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTrxWallet = `-- name: UpdateTrxWallet :one
UPDATE trx_wallets
SET wallet_id=$2, 
    amount=$3, 
    balance_before=$4, 
    balance=$5, 
    dest_model=$6, 
    dest_model_id=$7, 
    trx_receives_id=$8, 
    debit=$9,
    updated_at=now(),
    updated_by=$10
WHERE id = $1
RETURNING id, wallet_id, amount, balance_before, balance, dest_model, dest_model_id, trx_receives_id, created_at, created_by, updated_at, updated_by, debit
`

type UpdateTrxWalletParams struct {
	ID            int32          `json:"id"`
	WalletID      int32          `json:"wallet_id"`
	Amount        string         `json:"amount"`
	BalanceBefore string         `json:"balance_before"`
	Balance       string         `json:"balance"`
	DestModel     sql.NullString `json:"dest_model"`
	DestModelID   sql.NullInt32  `json:"dest_model_id"`
	TrxReceivesID int32          `json:"trx_receives_id"`
	Debit         bool           `json:"debit"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) UpdateTrxWallet(ctx context.Context, arg UpdateTrxWalletParams) (TrxWallet, error) {
	row := q.db.QueryRowContext(ctx, updateTrxWallet,
		arg.ID,
		arg.WalletID,
		arg.Amount,
		arg.BalanceBefore,
		arg.Balance,
		arg.DestModel,
		arg.DestModelID,
		arg.TrxReceivesID,
		arg.Debit,
		arg.UpdatedBy,
	)
	var i TrxWallet
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.BalanceBefore,
		&i.Balance,
		&i.DestModel,
		&i.DestModelID,
		&i.TrxReceivesID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Debit,
	)
	return i, err
}