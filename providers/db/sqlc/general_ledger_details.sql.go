// Code generated by sqlc. DO NOT EDIT.
// source: general_ledger_details.sql

package db

import (
	"context"
)

const createGeneralLedgerDetail = `-- name: CreateGeneralLedgerDetail :one
INSERT INTO general_ledger_details (
    owner_id, 
    general_ledger_id, 
    account_id, 
    debit, 
    credit, 
    balance_before, 
    balance, 
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, owner_id, general_ledger_id, account_id, debit, credit, balance_before, balance, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type CreateGeneralLedgerDetailParams struct {
	OwnerID         int32  `json:"owner_id"`
	GeneralLedgerID int32  `json:"general_ledger_id"`
	AccountID       int32  `json:"account_id"`
	Debit           string `json:"debit"`
	Credit          string `json:"credit"`
	BalanceBefore   string `json:"balance_before"`
	Balance         string `json:"balance"`
	CreatedBy       int32  `json:"created_by"`
	UpdatedBy       int32  `json:"updated_by"`
}

func (q *Queries) CreateGeneralLedgerDetail(ctx context.Context, arg CreateGeneralLedgerDetailParams) (GeneralLedgerDetail, error) {
	row := q.db.QueryRowContext(ctx, createGeneralLedgerDetail,
		arg.OwnerID,
		arg.GeneralLedgerID,
		arg.AccountID,
		arg.Debit,
		arg.Credit,
		arg.BalanceBefore,
		arg.Balance,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i GeneralLedgerDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.GeneralLedgerID,
		&i.AccountID,
		&i.Debit,
		&i.Credit,
		&i.BalanceBefore,
		&i.Balance,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const deleteGeneralLedgerDetail = `-- name: DeleteGeneralLedgerDetail :exec
DELETE FROM general_ledger_details
WHERE id = $1
`

func (q *Queries) DeleteGeneralLedgerDetail(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteGeneralLedgerDetail, id)
	return err
}

const getGeneralLedgerDetail = `-- name: GetGeneralLedgerDetail :one
SELECT id, owner_id, general_ledger_id, account_id, debit, credit, balance_before, balance, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM general_ledger_details
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGeneralLedgerDetail(ctx context.Context, id int32) (GeneralLedgerDetail, error) {
	row := q.db.QueryRowContext(ctx, getGeneralLedgerDetail, id)
	var i GeneralLedgerDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.GeneralLedgerID,
		&i.AccountID,
		&i.Debit,
		&i.Credit,
		&i.BalanceBefore,
		&i.Balance,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getGeneralLedgerDetailForUpdate = `-- name: GetGeneralLedgerDetailForUpdate :one
SELECT id, owner_id, general_ledger_id, account_id, debit, credit, balance_before, balance, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM general_ledger_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetGeneralLedgerDetailForUpdate(ctx context.Context, id int32) (GeneralLedgerDetail, error) {
	row := q.db.QueryRowContext(ctx, getGeneralLedgerDetailForUpdate, id)
	var i GeneralLedgerDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.GeneralLedgerID,
		&i.AccountID,
		&i.Debit,
		&i.Credit,
		&i.BalanceBefore,
		&i.Balance,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const listGeneralLedgerDetails = `-- name: ListGeneralLedgerDetails :many
SELECT id, owner_id, general_ledger_id, account_id, debit, credit, balance_before, balance, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM general_ledger_details
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListGeneralLedgerDetailsParams struct {
	OwnerID int32 `json:"owner_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListGeneralLedgerDetails(ctx context.Context, arg ListGeneralLedgerDetailsParams) ([]GeneralLedgerDetail, error) {
	rows, err := q.db.QueryContext(ctx, listGeneralLedgerDetails, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GeneralLedgerDetail
	for rows.Next() {
		var i GeneralLedgerDetail
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.GeneralLedgerID,
			&i.AccountID,
			&i.Debit,
			&i.Credit,
			&i.BalanceBefore,
			&i.Balance,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
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

const updateGeneralLedgerDetail = `-- name: UpdateGeneralLedgerDetail :one
UPDATE general_ledger_details
SET owner_id=$2, 
    general_ledger_id=$3, 
    account_id=$4, 
    debit=$5, 
    credit=$6, 
    balance_before=$7, 
    balance=$8, 
    updated_at=now(),
    updated_by=$9
WHERE id = $1
RETURNING id, owner_id, general_ledger_id, account_id, debit, credit, balance_before, balance, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type UpdateGeneralLedgerDetailParams struct {
	ID              int32  `json:"id"`
	OwnerID         int32  `json:"owner_id"`
	GeneralLedgerID int32  `json:"general_ledger_id"`
	AccountID       int32  `json:"account_id"`
	Debit           string `json:"debit"`
	Credit          string `json:"credit"`
	BalanceBefore   string `json:"balance_before"`
	Balance         string `json:"balance"`
	UpdatedBy       int32  `json:"updated_by"`
}

func (q *Queries) UpdateGeneralLedgerDetail(ctx context.Context, arg UpdateGeneralLedgerDetailParams) (GeneralLedgerDetail, error) {
	row := q.db.QueryRowContext(ctx, updateGeneralLedgerDetail,
		arg.ID,
		arg.OwnerID,
		arg.GeneralLedgerID,
		arg.AccountID,
		arg.Debit,
		arg.Credit,
		arg.BalanceBefore,
		arg.Balance,
		arg.UpdatedBy,
	)
	var i GeneralLedgerDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.GeneralLedgerID,
		&i.AccountID,
		&i.Debit,
		&i.Credit,
		&i.BalanceBefore,
		&i.Balance,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}