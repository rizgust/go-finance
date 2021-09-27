INSERT INTO general_ledger_details (
    onwer_id, 
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

-- name: GetGeneralLedgerDetail :one
SELECT * FROM general_ledger_details
WHERE id = $1 LIMIT 1;

-- name: GetGeneralLedgerDetailForUpdate :one
SELECT * FROM general_ledger_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListGeneralLedgerDetails :many
SELECT * FROM general_ledger_details
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateGeneralLedgerDetail :one
UPDATE general_ledger_details
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteGeneralLedgerDetail :exec
DELETE FROM general_ledger_details
WHERE id = $1;
