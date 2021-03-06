// Code generated by sqlc. DO NOT EDIT.
// source: journal_details.sql

package db

import (
	"context"
	"database/sql"
)

const createJournalDetail = `-- name: CreateJournalDetail :one
INSERT INTO journal_details (
    owner_id, 
    account_id, 
    journal_id, 
    amount, 
    reff_model,
    reff_model_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, owner_id, account_id, journal_id, amount, reff_model, reff_model_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type CreateJournalDetailParams struct {
	OwnerID     int32          `json:"owner_id"`
	AccountID   int64          `json:"account_id"`
	JournalID   int32          `json:"journal_id"`
	Amount      string         `json:"amount"`
	ReffModel   sql.NullString `json:"reff_model"`
	ReffModelID sql.NullInt32  `json:"reff_model_id"`
	CreatedBy   int32          `json:"created_by"`
	UpdatedBy   int32          `json:"updated_by"`
}

func (q *Queries) CreateJournalDetail(ctx context.Context, arg CreateJournalDetailParams) (JournalDetail, error) {
	row := q.db.QueryRowContext(ctx, createJournalDetail,
		arg.OwnerID,
		arg.AccountID,
		arg.JournalID,
		arg.Amount,
		arg.ReffModel,
		arg.ReffModelID,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i JournalDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.AccountID,
		&i.JournalID,
		&i.Amount,
		&i.ReffModel,
		&i.ReffModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const deleteJournalDetail = `-- name: DeleteJournalDetail :exec
UPDATE journal_details
SET deleted_at=now(),
    deleted_by=$2
WHERE id = $1
RETURNING id, owner_id, account_id, journal_id, amount, reff_model, reff_model_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type DeleteJournalDetailParams struct {
	ID        int32         `json:"id"`
	DeletedBy sql.NullInt32 `json:"deleted_by"`
}

func (q *Queries) DeleteJournalDetail(ctx context.Context, arg DeleteJournalDetailParams) error {
	_, err := q.db.ExecContext(ctx, deleteJournalDetail, arg.ID, arg.DeletedBy)
	return err
}

const getJournalDetail = `-- name: GetJournalDetail :one
SELECT id, owner_id, account_id, journal_id, amount, reff_model, reff_model_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM journal_details
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetJournalDetail(ctx context.Context, id int32) (JournalDetail, error) {
	row := q.db.QueryRowContext(ctx, getJournalDetail, id)
	var i JournalDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.AccountID,
		&i.JournalID,
		&i.Amount,
		&i.ReffModel,
		&i.ReffModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getJournalDetailForUpdate = `-- name: GetJournalDetailForUpdate :one
SELECT id, owner_id, account_id, journal_id, amount, reff_model, reff_model_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM journal_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetJournalDetailForUpdate(ctx context.Context, id int32) (JournalDetail, error) {
	row := q.db.QueryRowContext(ctx, getJournalDetailForUpdate, id)
	var i JournalDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.AccountID,
		&i.JournalID,
		&i.Amount,
		&i.ReffModel,
		&i.ReffModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const listJournalDetails = `-- name: ListJournalDetails :many
SELECT id, owner_id, account_id, journal_id, amount, reff_model, reff_model_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM journal_details
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListJournalDetailsParams struct {
	OwnerID int32 `json:"owner_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListJournalDetails(ctx context.Context, arg ListJournalDetailsParams) ([]JournalDetail, error) {
	rows, err := q.db.QueryContext(ctx, listJournalDetails, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []JournalDetail
	for rows.Next() {
		var i JournalDetail
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.AccountID,
			&i.JournalID,
			&i.Amount,
			&i.ReffModel,
			&i.ReffModelID,
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

const updateJournalDetail = `-- name: UpdateJournalDetail :one
UPDATE journal_details
SET owner_id=$2, 
    account_id=$3, 
    journal_id=$4, 
    amount=$5, 
    reff_model=$6,
    reff_model_id=$7,
    updated_at=now(),
    updated_by=$8
WHERE id = $1
RETURNING id, owner_id, account_id, journal_id, amount, reff_model, reff_model_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type UpdateJournalDetailParams struct {
	ID          int32          `json:"id"`
	OwnerID     int32          `json:"owner_id"`
	AccountID   int64          `json:"account_id"`
	JournalID   int32          `json:"journal_id"`
	Amount      string         `json:"amount"`
	ReffModel   sql.NullString `json:"reff_model"`
	ReffModelID sql.NullInt32  `json:"reff_model_id"`
	UpdatedBy   int32          `json:"updated_by"`
}

func (q *Queries) UpdateJournalDetail(ctx context.Context, arg UpdateJournalDetailParams) (JournalDetail, error) {
	row := q.db.QueryRowContext(ctx, updateJournalDetail,
		arg.ID,
		arg.OwnerID,
		arg.AccountID,
		arg.JournalID,
		arg.Amount,
		arg.ReffModel,
		arg.ReffModelID,
		arg.UpdatedBy,
	)
	var i JournalDetail
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.AccountID,
		&i.JournalID,
		&i.Amount,
		&i.ReffModel,
		&i.ReffModelID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}
