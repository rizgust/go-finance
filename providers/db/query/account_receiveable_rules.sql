INSERT INTO account_receiveable_rules (
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

-- name: GetAccountReceiveableRule :one
SELECT * FROM account_receiveable_rules
WHERE id = $1 LIMIT 1;

-- name: GetAccountReceiveableRuleForUpdate :one
SELECT * FROM account_receiveable_rules
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccountReceiveableRules :many
SELECT * FROM account_receiveable_rules
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccountReceiveableRule :one
UPDATE account_receiveable_rules
SET owner_id=$2, 
    ar_id=$3, 
    period_id=$4, 
    "rule"=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteAccountReceiveableRule :exec
DELETE FROM account_receiveable_rules
WHERE id = $1;
