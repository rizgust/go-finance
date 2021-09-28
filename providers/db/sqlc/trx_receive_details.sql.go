// Code generated by sqlc. DO NOT EDIT.
// source: trx_receive_details.sql

package db

import (
	"context"
	"database/sql"
)

const createTrxReceiveDetail = `-- name: CreateTrxReceiveDetail :one
INSERT INTO trx_receive_details (
    trx_receives_id, 
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
) RETURNING id, trx_receives_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by
`

type CreateTrxReceiveDetailParams struct {
	TrxReceivesID int32          `json:"trx_receives_id"`
	Amount        string         `json:"amount"`
	IsSource      bool           `json:"is_source"`
	AccountID     int32          `json:"account_id"`
	Model         sql.NullString `json:"model"`
	ModelID       sql.NullInt32  `json:"model_id"`
	CreatedBy     int32          `json:"created_by"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) CreateTrxReceiveDetail(ctx context.Context, arg CreateTrxReceiveDetailParams) (TrxReceiveDetail, error) {
	row := q.db.QueryRowContext(ctx, createTrxReceiveDetail,
		arg.TrxReceivesID,
		arg.Amount,
		arg.IsSource,
		arg.AccountID,
		arg.Model,
		arg.ModelID,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i TrxReceiveDetail
	err := row.Scan(
		&i.ID,
		&i.TrxReceivesID,
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

const deleteTrxReceiveDetail = `-- name: DeleteTrxReceiveDetail :exec
DELETE FROM trx_receive_details
WHERE id = $1
`

func (q *Queries) DeleteTrxReceiveDetail(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTrxReceiveDetail, id)
	return err
}

const getTrxReceiveDetail = `-- name: GetTrxReceiveDetail :one
SELECT id, trx_receives_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by FROM trx_receive_details
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTrxReceiveDetail(ctx context.Context, id int32) (TrxReceiveDetail, error) {
	row := q.db.QueryRowContext(ctx, getTrxReceiveDetail, id)
	var i TrxReceiveDetail
	err := row.Scan(
		&i.ID,
		&i.TrxReceivesID,
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

const getTrxReceiveDetailForUpdate = `-- name: GetTrxReceiveDetailForUpdate :one
SELECT id, trx_receives_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by FROM trx_receive_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTrxReceiveDetailForUpdate(ctx context.Context, id int32) (TrxReceiveDetail, error) {
	row := q.db.QueryRowContext(ctx, getTrxReceiveDetailForUpdate, id)
	var i TrxReceiveDetail
	err := row.Scan(
		&i.ID,
		&i.TrxReceivesID,
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

const listTrxReceiveDetails = `-- name: ListTrxReceiveDetails :many
SELECT id, trx_receives_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by FROM trx_receive_details
WHERE trx_receives_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTrxReceiveDetailsParams struct {
	TrxReceivesID int32 `json:"trx_receives_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTrxReceiveDetails(ctx context.Context, arg ListTrxReceiveDetailsParams) ([]TrxReceiveDetail, error) {
	rows, err := q.db.QueryContext(ctx, listTrxReceiveDetails, arg.TrxReceivesID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TrxReceiveDetail
	for rows.Next() {
		var i TrxReceiveDetail
		if err := rows.Scan(
			&i.ID,
			&i.TrxReceivesID,
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

const updateTrxReceiveDetail = `-- name: UpdateTrxReceiveDetail :one
UPDATE trx_receive_details
SET trx_receives_id=$2, 
    amount=$3, 
    is_source=$4, 
    account_id=$5, 
    model=$6, 
    model_id=$7,
    updated_at=now(),
    updated_by=$8
WHERE id = $1
RETURNING id, trx_receives_id, amount, is_source, account_id, model, model_id, created_at, created_by, updated_at, updated_by
`

type UpdateTrxReceiveDetailParams struct {
	ID            int32          `json:"id"`
	TrxReceivesID int32          `json:"trx_receives_id"`
	Amount        string         `json:"amount"`
	IsSource      bool           `json:"is_source"`
	AccountID     int32          `json:"account_id"`
	Model         sql.NullString `json:"model"`
	ModelID       sql.NullInt32  `json:"model_id"`
	UpdatedBy     int32          `json:"updated_by"`
}

func (q *Queries) UpdateTrxReceiveDetail(ctx context.Context, arg UpdateTrxReceiveDetailParams) (TrxReceiveDetail, error) {
	row := q.db.QueryRowContext(ctx, updateTrxReceiveDetail,
		arg.ID,
		arg.TrxReceivesID,
		arg.Amount,
		arg.IsSource,
		arg.AccountID,
		arg.Model,
		arg.ModelID,
		arg.UpdatedBy,
	)
	var i TrxReceiveDetail
	err := row.Scan(
		&i.ID,
		&i.TrxReceivesID,
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
