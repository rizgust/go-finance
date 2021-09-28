-- name: CreateMBankAccount :one
INSERT INTO m_bank_accounts (
    owner_id, 
    bank_id, 
    branch_name, 
    "name", 
    "number", 
    created_by, 
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetMBankAccount :one
SELECT * FROM m_bank_accounts
WHERE id = $1 LIMIT 1;

-- name: GetMBankAccountForUpdate :one
SELECT * FROM m_bank_accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMBankAccounts :many
SELECT * FROM m_bank_accounts
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateMBankAccount :one
UPDATE m_bank_accounts
SET owner_id=$2, 
    bank_id=$3, 
    branch_name=$4, 
    "name"=$5, 
    "number"=$6, 
    updated_at=now(),
    updated_by=$7
WHERE id = $1
RETURNING *;

-- name: DeleteMBankAccount :exec
DELETE FROM m_bank_accounts
WHERE id = $1;