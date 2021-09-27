INSERT INTO trx_receive_details (
    trx_receives_id, 
    amount, 
    is_source, 
    account_id, 
    model, 
    model_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetTrxReceiveDetail :one
SELECT * FROM trx_receive_details
WHERE id = $1 LIMIT 1;

-- name: GetTrxReceiveDetailForUpdate :one
SELECT * FROM trx_receive_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrxReceiveDetails :many
SELECT * FROM trx_receive_details
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTrxReceiveDetail :one
UPDATE trx_receive_details
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteTrxReceiveDetail :exec
DELETE FROM trx_receive_details
WHERE id = $1;
