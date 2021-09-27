INSERT INTO wallets (
    owner_id, 
    user_id, 
    code, 
    "name", 
    alias, 
    balance, 
    account_id, 
    other_info, 
    allow_minus,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1;

-- name: GetWalletForUpdate :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListWallets :many
SELECT * FROM wallets
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateWallet :one
UPDATE wallets
SET owner_id=$2, 
    user_id=$3, 
    code=$4, 
    "name"=$5, 
    alias=$6, 
    balance=$7, 
    account_id=$8, 
    other_info=$9, 
    allow_minus=$10,
    updated_at=now(),
    updated_by=$11
WHERE id = $1
RETURNING *;

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1;
