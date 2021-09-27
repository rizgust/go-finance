INSERT INTO m_banks (
    owner_id, 
    code, 
    "name", 
    alias,  
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetMBank :one
SELECT * FROM m_banks
WHERE id = $1 LIMIT 1;

-- name: GetMBankForUpdate :one
SELECT * FROM m_banks
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMBanks :many
SELECT * FROM m_banks
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateMBank :one
UPDATE m_banks
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteMBank :exec
DELETE FROM m_banks
WHERE id = $1;
