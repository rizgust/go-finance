// Code generated by sqlc. DO NOT EDIT.
// source: account_payables.sql

package db

import (
	"context"
	"database/sql"
)

const createAccountPayable = `-- name: CreateAccountPayable :one
INSERT INTO account_payables (
    owner_id, 
    code, 
    "name", 
    description, 
    src_account_id, 
    dst_account_id, 
    recurring_type, 
    recurring_periode, 
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING id, owner_id, code, name, description, src_account_id, dst_account_id, recurring_type, recurring_periode, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type CreateAccountPayableParams struct {
	OwnerID          int32          `json:"owner_id"`
	Code             string         `json:"code"`
	Name             string         `json:"name"`
	Description      sql.NullString `json:"description"`
	SrcAccountID     int32          `json:"src_account_id"`
	DstAccountID     int32          `json:"dst_account_id"`
	RecurringType    int16          `json:"recurring_type"`
	RecurringPeriode int16          `json:"recurring_periode"`
	CreatedBy        int32          `json:"created_by"`
	UpdatedBy        int32          `json:"updated_by"`
}

func (q *Queries) CreateAccountPayable(ctx context.Context, arg CreateAccountPayableParams) (AccountPayable, error) {
	row := q.db.QueryRowContext(ctx, createAccountPayable,
		arg.OwnerID,
		arg.Code,
		arg.Name,
		arg.Description,
		arg.SrcAccountID,
		arg.DstAccountID,
		arg.RecurringType,
		arg.RecurringPeriode,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i AccountPayable
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.SrcAccountID,
		&i.DstAccountID,
		&i.RecurringType,
		&i.RecurringPeriode,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const deleteAccountPayable = `-- name: DeleteAccountPayable :exec
DELETE FROM account_payables
WHERE id = $1
`

func (q *Queries) DeleteAccountPayable(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccountPayable, id)
	return err
}

const getAccountPayable = `-- name: GetAccountPayable :one
SELECT id, owner_id, code, name, description, src_account_id, dst_account_id, recurring_type, recurring_periode, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM account_payables
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccountPayable(ctx context.Context, id int32) (AccountPayable, error) {
	row := q.db.QueryRowContext(ctx, getAccountPayable, id)
	var i AccountPayable
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.SrcAccountID,
		&i.DstAccountID,
		&i.RecurringType,
		&i.RecurringPeriode,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getAccountPayableForUpdate = `-- name: GetAccountPayableForUpdate :one
SELECT id, owner_id, code, name, description, src_account_id, dst_account_id, recurring_type, recurring_periode, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM account_payables
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountPayableForUpdate(ctx context.Context, id int32) (AccountPayable, error) {
	row := q.db.QueryRowContext(ctx, getAccountPayableForUpdate, id)
	var i AccountPayable
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.SrcAccountID,
		&i.DstAccountID,
		&i.RecurringType,
		&i.RecurringPeriode,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const listAccountPayables = `-- name: ListAccountPayables :many
SELECT id, owner_id, code, name, description, src_account_id, dst_account_id, recurring_type, recurring_periode, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM account_payables
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListAccountPayablesParams struct {
	OwnerID int32 `json:"owner_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListAccountPayables(ctx context.Context, arg ListAccountPayablesParams) ([]AccountPayable, error) {
	rows, err := q.db.QueryContext(ctx, listAccountPayables, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccountPayable
	for rows.Next() {
		var i AccountPayable
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Code,
			&i.Name,
			&i.Description,
			&i.SrcAccountID,
			&i.DstAccountID,
			&i.RecurringType,
			&i.RecurringPeriode,
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

const updateAccountPayable = `-- name: UpdateAccountPayable :one
UPDATE account_payables
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    description=$5, 
    src_account_id=$6, 
    dst_account_id=$7, 
    recurring_type=$8, 
    recurring_periode=$9, 
    updated_at=now(),
    updated_by=$10
WHERE id = $1
RETURNING id, owner_id, code, name, description, src_account_id, dst_account_id, recurring_type, recurring_periode, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type UpdateAccountPayableParams struct {
	ID               int32          `json:"id"`
	OwnerID          int32          `json:"owner_id"`
	Code             string         `json:"code"`
	Name             string         `json:"name"`
	Description      sql.NullString `json:"description"`
	SrcAccountID     int32          `json:"src_account_id"`
	DstAccountID     int32          `json:"dst_account_id"`
	RecurringType    int16          `json:"recurring_type"`
	RecurringPeriode int16          `json:"recurring_periode"`
	UpdatedBy        int32          `json:"updated_by"`
}

func (q *Queries) UpdateAccountPayable(ctx context.Context, arg UpdateAccountPayableParams) (AccountPayable, error) {
	row := q.db.QueryRowContext(ctx, updateAccountPayable,
		arg.ID,
		arg.OwnerID,
		arg.Code,
		arg.Name,
		arg.Description,
		arg.SrcAccountID,
		arg.DstAccountID,
		arg.RecurringType,
		arg.RecurringPeriode,
		arg.UpdatedBy,
	)
	var i AccountPayable
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.SrcAccountID,
		&i.DstAccountID,
		&i.RecurringType,
		&i.RecurringPeriode,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}