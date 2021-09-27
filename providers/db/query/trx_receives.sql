INSERT INTO trx_receives (
    onwer_id, 
    user_id, 
    "number", 
    amount, 
    "date", 
    status, 
    description,
    payment_method,
    journal_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetTrxReceive :one
SELECT * FROM trx_receives
WHERE id = $1 LIMIT 1;

-- name: GetTrxReceiveForUpdate :one
SELECT * FROM trx_receives
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrxReceives :many
SELECT * FROM trx_receives
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTrxReceive :one
UPDATE trx_receives
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteTrxReceive :exec
DELETE FROM trx_receives
WHERE id = $1;
