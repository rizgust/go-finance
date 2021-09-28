-- name: CreateAccountPayableRule :one
INSERT INTO account_payable_rules (
    owner_id, 
    ap_id, 
    period_id,
    "rule",
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAccountPayableRule :one
SELECT * FROM account_payable_rules
WHERE id = $1 LIMIT 1;

-- name: GetAccountPayableRuleForUpdate :one
SELECT * FROM account_payable_rules
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccountPayableRules :many
SELECT * FROM account_payable_rules
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccountPayableRule :one
UPDATE account_payable_rules
SET owner_id=$2, 
    ap_id=$3, 
    period_id=$4, 
    "rule"=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteAccountPayableRule :exec
DELETE FROM account_payable_rules
WHERE id = $1;
