INSERT INTO trx_wallets (
    wallet_id, 
    amount, 
    balance_before, 
    balance, 
    dest_model, 
    dest_model_id, 
    trx_receives_id, 
    debit,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetTrxWallet :one
SELECT * FROM trx_wallets
WHERE id = $1 LIMIT 1;

-- name: GetTrxWalletForUpdate :one
SELECT * FROM trx_wallets
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrxWallets :many
SELECT * FROM trx_wallets
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTrxWallet :one
UPDATE trx_wallets
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteTrxWallet :exec
DELETE FROM trx_wallets
WHERE id = $1;
