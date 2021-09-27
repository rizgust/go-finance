INSERT INTO account_receivable_rules (
    owner_id, 
    ar_id, 
    period_id,
    "rule",
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAccountReceivableRule :one
SELECT * FROM account_receivable_rules
WHERE id = $1 LIMIT 1;

-- name: GetAccountReceivableRuleForUpdate :one
SELECT * FROM account_receivable_rules
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccountReceivableRules :many
SELECT * FROM account_receivable_rules
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccountReceivableRule :one
UPDATE account_receivable_rules
SET owner_id=$2, 
    ar_id=$3, 
    period_id=$4, 
    "rule"=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteAccountReceivableRule :exec
DELETE FROM account_receivable_rules
WHERE id = $1;
