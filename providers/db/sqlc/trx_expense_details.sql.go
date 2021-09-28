// Code generated by sqlc. DO NOT EDIT.
// source: trx_expense_details.sql

package db

import (
	"context"
	"database/sql"
)

const createTrxExpenseDetail = `-- name: CreateTrxExpenseDetail :one
INSERT INTO trx_expense_details (
    trx_expenses_id, 
    amount, 
    is_source, 
    account_id, 
    model, 
    model_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, trx_expenses_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by
`

type CreateTrxExpenseDetailParams struct {
	TrxExpensesID int32          `json:"trx_expenses_id"`
	Amount        string         `json:"amount"`
	IsSource      bool           `json:"is_source"`
	AccountID     int32          `json:"account_id"`
	Model         sql.NullString `json:"model"`
	ModelID       sql.NullInt32  `json:"model_id"`
	CreatedBy     int32          `json:"created_by"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) CreateTrxExpenseDetail(ctx context.Context, arg CreateTrxExpenseDetailParams) (TrxExpenseDetail, error) {
	row := q.db.QueryRowContext(ctx, createTrxExpenseDetail,
		arg.TrxExpensesID,
		arg.Amount,
		arg.IsSource,
		arg.AccountID,
		arg.Model,
		arg.ModelID,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i TrxExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.TrxExpensesID,
		&i.Amount,
		&i.IsSource,
		&i.AccountID,
		&i.Model,
		&i.ModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const deleteTrxExpenseDetail = `-- name: DeleteTrxExpenseDetail :exec
DELETE FROM trx_expense_details
WHERE id = $1
`

func (q *Queries) DeleteTrxExpenseDetail(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTrxExpenseDetail, id)
	return err
}

const getTrxExpenseDetail = `-- name: GetTrxExpenseDetail :one
SELECT id, trx_expenses_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by FROM trx_expense_details
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTrxExpenseDetail(ctx context.Context, id int32) (TrxExpenseDetail, error) {
	row := q.db.QueryRowContext(ctx, getTrxExpenseDetail, id)
	var i TrxExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.TrxExpensesID,
		&i.Amount,
		&i.IsSource,
		&i.AccountID,
		&i.Model,
		&i.ModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const getTrxExpenseDetailForUpdate = `-- name: GetTrxExpenseDetailForUpdate :one
SELECT id, trx_expenses_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by FROM trx_expense_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTrxExpenseDetailForUpdate(ctx context.Context, id int32) (TrxExpenseDetail, error) {
	row := q.db.QueryRowContext(ctx, getTrxExpenseDetailForUpdate, id)
	var i TrxExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.TrxExpensesID,
		&i.Amount,
		&i.IsSource,
		&i.AccountID,
		&i.Model,
		&i.ModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const listTrxExpenseDetails = `-- name: ListTrxExpenseDetails :many
SELECT id, trx_expenses_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by FROM trx_expense_details
WHERE trx_expenses_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTrxExpenseDetailsParams struct {
	TrxExpensesID int32 `json:"trx_expenses_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTrxExpenseDetails(ctx context.Context, arg ListTrxExpenseDetailsParams) ([]TrxExpenseDetail, error) {
	rows, err := q.db.QueryContext(ctx, listTrxExpenseDetails, arg.TrxExpensesID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TrxExpenseDetail
	for rows.Next() {
		var i TrxExpenseDetail
		if err := rows.Scan(
			&i.ID,
			&i.TrxExpensesID,
			&i.Amount,
			&i.IsSource,
			&i.AccountID,
			&i.Model,
			&i.ModelID,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
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

const updateTrxExpenseDetail = `-- name: UpdateTrxExpenseDetail :one
UPDATE trx_expense_details
SET trx_expenses_id=$2, 
    amount=$3, 
    is_source=$4, 
    account_id=$5, 
    model=$6, 
    model_id=$7,
    updated_at=now(),
    updated_by=$8
WHERE id = $1
RETURNING id, trx_expenses_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by
`

type UpdateTrxExpenseDetailParams struct {
	ID            int32          `json:"id"`
	TrxExpensesID int32          `json:"trx_expenses_id"`
	Amount        string         `json:"amount"`
	IsSource      bool           `json:"is_source"`
	AccountID     int32          `json:"account_id"`
	Model         sql.NullString `json:"model"`
	ModelID       sql.NullInt32  `json:"model_id"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) UpdateTrxExpenseDetail(ctx context.Context, arg UpdateTrxExpenseDetailParams) (TrxExpenseDetail, error) {
	row := q.db.QueryRowContext(ctx, updateTrxExpenseDetail,
		arg.ID,
		arg.TrxExpensesID,
		arg.Amount,
		arg.IsSource,
		arg.AccountID,
		arg.Model,
		arg.ModelID,
		arg.UpdatedBy,
	)
	var i TrxExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.TrxExpensesID,
		&i.Amount,
		&i.IsSource,
		&i.AccountID,
		&i.Model,
		&i.ModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}
