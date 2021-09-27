INSERT INTO account_receiveables (
    owner_id, 
    code, 
    "name", 
    description, 
    src_account_id, 
    dst_account_id, 
    recurring_type, 
    recurring_periode, 
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetAccountReceiveable :one
SELECT * FROM account_receiveables
WHERE id = $1 LIMIT 1;

-- name: GetAccountReceiveableForUpdate :one
SELECT * FROM account_receiveables
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccountReceiveables :many
SELECT * FROM account_receiveables
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccountReceiveable :one
UPDATE account_receiveables
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    description=$5, 
    src_account_id=$6, 
    dst_account_id=$7, 
    recurring_type=$8, 
    recurring_periode=$9, 
    updated_at=now(),
    updated_by=$10
WHERE id = $1
RETURNING *;

-- name: DeleteAccountReceiveable :exec
DELETE FROM account_receiveables
WHERE id = $1;
