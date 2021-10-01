-- name: CreateJournal :one
INSERT INTO journals (
    owner_id, 
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
    "number"=$3, 
    "date"=$4, 
    general_ledger_id=$5, 
    status=$6, 
    amount=$7, 
    description=$8, 
    reff_model=$9, 
    reff_model_id=$10,
    updated_at=now(),
    updated_by=$11
WHERE id = $1
RETURNING *;

-- name: DeleteJournal :exec
UPDATE journals
SET deleted_at=now(),
    deleted_by=$2
WHERE id = $1
RETURNING *;