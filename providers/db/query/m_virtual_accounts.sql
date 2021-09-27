INSERT INTO m_virtual_accounts (
    owner_id, 
    bank_id, 
    "number",
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetMVirtualAccount :one
SELECT * FROM m_virtual_accounts
WHERE id = $1 LIMIT 1;

-- name: GetMVirtualAccountForUpdate :one
SELECT * FROM m_virtual_accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMVirtualAccounts :many
SELECT * FROM m_virtual_accounts
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateMVirtualAccount :one
UPDATE m_virtual_accounts
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteMVirtualAccount :exec
DELETE FROM m_virtual_accounts
WHERE id = $1;
