// Code generated by sqlc. DO NOT EDIT.
// source: trx_expenses.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTrxExpense = `-- name: CreateTrxExpense :one
INSERT INTO trx_expenses (
    owner_id, 
    user_id, 
    "number", 
    amount, 
    "date", 
    status, 
    description, 
    payment_method,
    journal_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, owner_id, user_id, number, amount, date, status, description, payment_method, created_at, created_by, updated_at, updated_by, journal_id
`

type CreateTrxExpenseParams struct {
	OwnerID       int32          `json:"owner_id"`
	UserID        int32          `json:"user_id"`
	Number        sql.NullString `json:"number"`
	Amount        string         `json:"amount"`
	Date          time.Time      `json:"date"`
	Status        int16          `json:"status"`
	Description   sql.NullString `json:"description"`
	PaymentMethod int16          `json:"payment_method"`
	JournalID     sql.NullInt32  `json:"journal_id"`
	CreatedBy     int32          `json:"created_by"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) CreateTrxExpense(ctx context.Context, arg CreateTrxExpenseParams) (TrxExpense, error) {
	row := q.db.QueryRowContext(ctx, createTrxExpense,
		arg.OwnerID,
		arg.UserID,
		arg.Number,
		arg.Amount,
		arg.Date,
		arg.Status,
		arg.Description,
		arg.PaymentMethod,
		arg.JournalID,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i TrxExpense
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.UserID,
		&i.Number,
		&i.Amount,
		&i.Date,
		&i.Status,
		&i.Description,
		&i.PaymentMethod,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.JournalID,
	)
	return i, err
}

const deleteTrxExpense = `-- name: DeleteTrxExpense :exec
DELETE FROM trx_expenses
WHERE id = $1
`

func (q *Queries) DeleteTrxExpense(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTrxExpense, id)
	return err
}

const getTrxExpense = `-- name: GetTrxExpense :one
SELECT id, owner_id, user_id, number, amount, date, status, description, payment_method, created_at, created_by, updated_at, updated_by, journal_id FROM trx_expenses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTrxExpense(ctx context.Context, id int32) (TrxExpense, error) {
	row := q.db.QueryRowContext(ctx, getTrxExpense, id)
	var i TrxExpense
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.UserID,
		&i.Number,
		&i.Amount,
		&i.Date,
		&i.Status,
		&i.Description,
		&i.PaymentMethod,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.JournalID,
	)
	return i, err
}

const getTrxExpenseForUpdate = `-- name: GetTrxExpenseForUpdate :one
SELECT id, owner_id, user_id, number, amount, date, status, description, payment_method, created_at, created_by, updated_at, updated_by, journal_id FROM trx_expenses
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTrxExpenseForUpdate(ctx context.Context, id int32) (TrxExpense, error) {
	row := q.db.QueryRowContext(ctx, getTrxExpenseForUpdate, id)
	var i TrxExpense
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.UserID,
		&i.Number,
		&i.Amount,
		&i.Date,
		&i.Status,
		&i.Description,
		&i.PaymentMethod,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.JournalID,
	)
	return i, err
}

const listTrxExpenses = `-- name: ListTrxExpenses :many
SELECT id, owner_id, user_id, number, amount, date, status, description, payment_method, created_at, created_by, updated_at, updated_by, journal_id FROM trx_expenses
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTrxExpensesParams struct {
	OwnerID int32 `json:"owner_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListTrxExpenses(ctx context.Context, arg ListTrxExpensesParams) ([]TrxExpense, error) {
	rows, err := q.db.QueryContext(ctx, listTrxExpenses, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TrxExpense
	for rows.Next() {
		var i TrxExpense
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.UserID,
			&i.Number,
			&i.Amount,
			&i.Date,
			&i.Status,
			&i.Description,
			&i.PaymentMethod,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.JournalID,
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

const updateTrxExpense = `-- name: UpdateTrxExpense :one
UPDATE trx_expenses
SET owner_id=$2, 
    user_id=$3, 
    "number"=$4, 
    amount=$5, 
    "date"=$6, 
    status=$7, 
    description=$8, 
    payment_method=$9,
    journal_id=$10,
    updated_at=now(),
    updated_by=$11
WHERE id = $1
RETURNING id, owner_id, user_id, number, amount, date, status, description, payment_method, created_at, created_by, updated_at, updated_by, journal_id
`

type UpdateTrxExpenseParams struct {
	ID            int32          `json:"id"`
	OwnerID       int32          `json:"owner_id"`
	UserID        int32          `json:"user_id"`
	Number        sql.NullString `json:"number"`
	Amount        string         `json:"amount"`
	Date          time.Time      `json:"date"`
	Status        int16          `json:"status"`
	Description   sql.NullString `json:"description"`
	PaymentMethod int16          `json:"payment_method"`
	JournalID     sql.NullInt32  `json:"journal_id"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) UpdateTrxExpense(ctx context.Context, arg UpdateTrxExpenseParams) (TrxExpense, error) {
	row := q.db.QueryRowContext(ctx, updateTrxExpense,
		arg.ID,
		arg.OwnerID,
		arg.UserID,
		arg.Number,
		arg.Amount,
		arg.Date,
		arg.Status,
		arg.Description,
		arg.PaymentMethod,
		arg.JournalID,
		arg.UpdatedBy,
	)
	var i TrxExpense
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.UserID,
		&i.Number,
		&i.Amount,
		&i.Date,
		&i.Status,
		&i.Description,
		&i.PaymentMethod,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.JournalID,
	)
	return i, err
}
