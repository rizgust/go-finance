INSERT INTO journals (
    onwer_id, 
    "number", 
    "date", 
    general_ledger_id, 
    status, 
    amount, 
    description, 
    reff_model, 
    reff_model_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetJournal :one
SELECT * FROM journals
WHERE id = $1 LIMIT 1;

-- name: GetJournalForUpdate :one
SELECT * FROM journals
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListJournals :many
SELECT * FROM journals
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateJournal :one
UPDATE journals
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteJournal :exec
DELETE FROM journals
WHERE id = $1;
