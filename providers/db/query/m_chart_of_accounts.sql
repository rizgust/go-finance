-- name: CreateMChartOfAccount :one
INSERT INTO m_chart_of_accounts (
    owner_id, 
    code, 
    "name", 
    description, 
    category, 
    "level",
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetMChartOfAccount :one
SELECT * FROM m_chart_of_accounts
WHERE id = $1 LIMIT 1;

-- name: GetMChartOfAccountForUpdate :one
SELECT * FROM m_chart_of_accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMChartOfAccounts :many
SELECT * FROM m_chart_of_accounts
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateMChartOfAccount :one
UPDATE m_chart_of_accounts
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    description=$5, 
    category=$6, 
    "level"=$7,
    updated_at=now(),
    updated_by=$8
WHERE id = $1
RETURNING *;

-- name: DeleteMChartOfAccount :exec
UPDATE m_chart_of_accounts
SET deleted_at=now(),
    deleted_by=$2
WHERE id = $1
RETURNING *;