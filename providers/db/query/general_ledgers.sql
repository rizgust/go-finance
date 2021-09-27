INSERT INTO general_ledger_details (
    owner_id, 
    general_ledger_id, 
    account_id, 
    debit, 
    credit, 
    balance_before, 
    balance, 
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetGeneralLedger :one
SELECT * FROM general_ledger_details
WHERE id = $1 LIMIT 1;

-- name: GetGeneralLedgerForUpdate :one
SELECT * FROM general_ledger_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListGeneralLedgers :many
SELECT * FROM general_ledger_details
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateGeneralLedger :one
UPDATE general_ledger_details
SET owner_id=$2, 
    general_ledger_id=$3, 
    account_id=$4, 
    debit=$5, 
    credit=$6, 
    balance_before=$7, 
    balance=$8, 
    updated_at=now(),
    updated_by=$9
WHERE id = $1
RETURNING *;

-- name: DeleteGeneralLedger :exec
DELETE FROM general_ledger_details
WHERE id = $1;
