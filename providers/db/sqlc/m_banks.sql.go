// Code generated by sqlc. DO NOT EDIT.
// source: m_banks.sql

package db

import (
	"context"
	"database/sql"
)

const createMBank = `-- name: CreateMBank :one
INSERT INTO m_banks (
    owner_id, 
    code, 
    "name", 
    alias,  
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6
) RETURNING id, owner_id, code, name, alias, created_at, created_by, updated_at, updated_by
`

type CreateMBankParams struct {
	OwnerID   int32          `json:"owner_id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	Alias     sql.NullString `json:"alias"`
	CreatedBy int32          `json:"created_by"`
	UpdatedBy int32          `json:"updated_by"`
}

func (q *Queries) CreateMBank(ctx context.Context, arg CreateMBankParams) (MBank, error) {
	row := q.db.QueryRowContext(ctx, createMBank,
		arg.OwnerID,
		arg.Code,
		arg.Name,
		arg.Alias,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i MBank
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Alias,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const deleteMBank = `-- name: DeleteMBank :exec
DELETE FROM m_banks
WHERE id = $1
`

func (q *Queries) DeleteMBank(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteMBank, id)
	return err
}

const getMBank = `-- name: GetMBank :one
SELECT id, owner_id, code, name, alias, created_at, created_by, updated_at, updated_by FROM m_banks
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMBank(ctx context.Context, id int32) (MBank, error) {
	row := q.db.QueryRowContext(ctx, getMBank, id)
	var i MBank
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Alias,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const getMBankForUpdate = `-- name: GetMBankForUpdate :one
SELECT id, owner_id, code, name, alias, created_at, created_by, updated_at, updated_by FROM m_banks
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetMBankForUpdate(ctx context.Context, id int32) (MBank, error) {
	row := q.db.QueryRowContext(ctx, getMBankForUpdate, id)
	var i MBank
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Alias,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}

const listMBanks = `-- name: ListMBanks :many
SELECT id, owner_id, code, name, alias, created_at, created_by, updated_at, updated_by FROM m_banks
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListMBanksParams struct {
	OwnerID int32 `json:"owner_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListMBanks(ctx context.Context, arg ListMBanksParams) ([]MBank, error) {
	rows, err := q.db.QueryContext(ctx, listMBanks, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MBank
	for rows.Next() {
		var i MBank
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Code,
			&i.Name,
			&i.Alias,
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

const updateMBank = `-- name: UpdateMBank :one
UPDATE m_banks
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING id, owner_id, code, name, alias, created_at, created_by, updated_at, updated_by
`

type UpdateMBankParams struct {
	ID        int32          `json:"id"`
	OwnerID   int32          `json:"owner_id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	Alias     sql.NullString `json:"alias"`
	UpdatedBy int32          `json:"updated_by"`
}

func (q *Queries) UpdateMBank(ctx context.Context, arg UpdateMBankParams) (MBank, error) {
	row := q.db.QueryRowContext(ctx, updateMBank,
		arg.ID,
		arg.OwnerID,
		arg.Code,
		arg.Name,
		arg.Alias,
		arg.UpdatedBy,
	)
	var i MBank
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Code,
		&i.Name,
		&i.Alias,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	return i, err
}
