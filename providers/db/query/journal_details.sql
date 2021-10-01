-- name: CreateJournalDetail :one
INSERT INTO journal_details (
    owner_id, 
    account_id, 
    journal_id, 
    amount, 
    reff_model,
    reff_model_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetJournalDetail :one
SELECT * FROM journal_details
WHERE id = $1 LIMIT 1;

-- name: GetJournalDetailForUpdate :one
SELECT * FROM journal_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListJournalDetails :many
SELECT * FROM journal_details
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateJournalDetail :one
UPDATE journal_details
SET owner_id=$2, 
    account_id=$3, 
    journal_id=$4, 
    amount=$5, 
    reff_model=$6,
    reff_model_id=$7,
    updated_at=now(),
    updated_by=$8
WHERE id = $1
RETURNING *;

-- name: DeleteJournalDetail :exec
UPDATE journal_details
SET deleted_at=now(),
    deleted_by=$2
WHERE id = $1
RETURNING *;