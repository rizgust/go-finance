INSERT INTO trx_bank_transfers (
    ownerr_id, 
    user_id, 
    bank_id, 
    "number", 
    "date", 
    status, 
    amount, 
    description, 
    dest_model, 
    dest_model_id, 
    account_owner, 
    account_number, 
    additional_info, 
    trx_receives_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
) RETURNING *;

-- name: GetTrxBankTransfer :one
SELECT * FROM trx_bank_transfers
WHERE id = $1 LIMIT 1;

-- name: GetTrxBankTransferForUpdate :one
SELECT * FROM trx_bank_transfers
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrxBankTransfers :many
SELECT * FROM trx_bank_transfers
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTrxBankTransfer :one
UPDATE trx_bank_transfers
SET owner_id=$2, 
    user_id=$3, 
    bank_id=$4, 
    "number"=$5, 
    "date"=$6, 
    status=$7, 
    amount=$8, 
    description=$9, 
    dest_model=$10, 
    dest_model_id=$11, 
    account_owner=$12, 
    account_number=$13, 
    additional_info=$14, 
    trx_receives_id=$15,
    updated_at=now(),
    updated_by=$16
WHERE id = $1
RETURNING *;

-- name: DeleteTrxBankTransfer :exec
DELETE FROM trx_bank_transfers
WHERE id = $1;
