INSERT INTO m_discounts (
    onwer_id, 
    code, 
    "name", 
    description, 
    is_active, 
    value, 
    is_percent, 
    start_date, 
    end_date,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetMDiscount :one
SELECT * FROM m_discounts
WHERE id = $1 LIMIT 1;

-- name: GetMDiscountForUpdate :one
SELECT * FROM m_discounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMDiscounts :many
SELECT * FROM m_discounts
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateMDiscount :one
UPDATE m_discounts
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    description=$5, 
    is_active=$6, 
    value=$7, 
    is_percent=$8, 
    start_date=$9, 
    end_date=$10,
    updated_at=now(),
    updated_by=$11
WHERE id = $1
RETURNING *;

-- name: DeleteMDiscount :exec
DELETE FROM m_discounts
WHERE id = $1;
