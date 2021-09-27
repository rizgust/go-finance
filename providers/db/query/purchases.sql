INSERT INTO purchases (
    onwer_id, 
    user_id, 
    "number", 
    ap_id, 
    status, 
    amount, 
    amount_paid, 
    "date", 
    due_date, 
    additional_info, 
    period_id, 
    expense_id, 
    discount_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING *;

-- name: GetPurchase :one
SELECT * FROM purchases
WHERE id = $1 LIMIT 1;

-- name: GetPurchaseForUpdate :one
SELECT * FROM purchases
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListPurchases :many
SELECT * FROM purchases
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdatePurchase :one
UPDATE purchases
SET owner_id=$2,  
    user_id=$3, 
    "number"=$4, 
    ar_id=$5, 
    status=$6, 
    amount=$7, 
    amount_paid=$8, 
    "date"=$9, 
    due_date=$10, 
    additional_info=$11, 
    period_id=$12, 
    expense_id=$13, 
    discount_id=$14,
    updated_at=now(),
    updated_by=$15
WHERE id = $1
RETURNING *;

-- name: DeletePurchase :exec
DELETE FROM purchases
WHERE id = $1;
