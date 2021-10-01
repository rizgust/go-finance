// Code generated by sqlc. DO NOT EDIT.
// source: m_book_periods.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMBookPeriod = `-- name: CreateMBookPeriod :one
INSERT INTO m_book_periods (
    owner_id, 
    code, 
    "name", 
    description, 
    status, 
    start_date, 
    end_date,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, owner_id, code, name, description, status, start_date, end_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type CreateMBookPeriodParams struct {
	OwnerID     int32          `json:"owner_id"`
	Code        sql.NullString `json:"code"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Status      int16          `json:"status"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	CreatedBy   int32          `json:"created_by"`
	UpdatedBy   int32          `json:"updated_by"`
}

func (q *Queries) CreateMBookPeriod(ctx context.Context, arg CreateMBookPeriodParams) (MBookPeriod, error) {
	row := q.db.QueryRowContext(ctx, createMBookPeriod,
		arg.OwnerID,
		arg.Code,
		arg.Name,
		arg.Description,
		arg.Status,
		arg.StartDate,
		arg.EndDate,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i MBookPeriod
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const deleteMBookPeriod = `-- name: DeleteMBookPeriod :exec
UPDATE m_book_periods
SET deleted_at=now(),
    deleted_by=$2
WHERE id = $1
RETURNING id, owner_id, code, name, description, status, start_date, end_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type DeleteMBookPeriodParams struct {
	ID        int32         `json:"id"`
	DeletedBy sql.NullInt32 `json:"deleted_by"`
}

func (q *Queries) DeleteMBookPeriod(ctx context.Context, arg DeleteMBookPeriodParams) error {
	_, err := q.db.ExecContext(ctx, deleteMBookPeriod, arg.ID, arg.DeletedBy)
	return err
}

const getMBookPeriod = `-- name: GetMBookPeriod :one
SELECT id, owner_id, code, name, description, status, start_date, end_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM m_book_periods
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMBookPeriod(ctx context.Context, id int32) (MBookPeriod, error) {
	row := q.db.QueryRowContext(ctx, getMBookPeriod, id)
	var i MBookPeriod
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getMBookPeriodForUpdate = `-- name: GetMBookPeriodForUpdate :one
SELECT id, owner_id, code, name, description, status, start_date, end_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM m_book_periods
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetMBookPeriodForUpdate(ctx context.Context, id int32) (MBookPeriod, error) {
	row := q.db.QueryRowContext(ctx, getMBookPeriodForUpdate, id)
	var i MBookPeriod
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const listMBookPeriods = `-- name: ListMBookPeriods :many
SELECT id, owner_id, code, name, description, status, start_date, end_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM m_book_periods
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListMBookPeriodsParams struct {
	OwnerID int32 `json:"owner_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListMBookPeriods(ctx context.Context, arg ListMBookPeriodsParams) ([]MBookPeriod, error) {
	rows, err := q.db.QueryContext(ctx, listMBookPeriods, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MBookPeriod
	for rows.Next() {
		var i MBookPeriod
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Code,
			&i.Name,
			&i.Description,
			&i.Status,
			&i.StartDate,
			&i.EndDate,
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

const updateMBookPeriod = `-- name: UpdateMBookPeriod :one
UPDATE m_book_periods
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    description=$5, 
    status=$6, 
    start_date=$7, 
    end_date=$8, 
    updated_at=now(),
    updated_by=$9
WHERE id = $1
RETURNING id, owner_id, code, name, description, status, start_date, end_date, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

type UpdateMBookPeriodParams struct {
	ID          int32          `json:"id"`
	OwnerID     int32          `json:"owner_id"`
	Code        sql.NullString `json:"code"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Status      int16          `json:"status"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	UpdatedBy   int32          `json:"updated_by"`
}

func (q *Queries) UpdateMBookPeriod(ctx context.Context, arg UpdateMBookPeriodParams) (MBookPeriod, error) {
	row := q.db.QueryRowContext(ctx, updateMBookPeriod,
		arg.ID,
		arg.OwnerID,
		arg.Code,
		arg.Name,
		arg.Description,
		arg.Status,
		arg.StartDate,
		arg.EndDate,
		arg.UpdatedBy,
	)
	var i MBookPeriod
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}
