INSERT INTO trx_expenses (
    ownerr_id, 
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

-- name: GetTrxExpense :one
SELECT * FROM trx_expenses
WHERE id = $1 LIMIT 1;

-- name: GetTrxExpenseForUpdate :one
SELECT * FROM trx_expenses
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrxExpenses :many
SELECT * FROM trx_expenses
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTrxExpense :one
UPDATE trx_expenses
SET ownerr_id=$2, 
    user_id=$3, 
    "number"=$4, 
    amount=$5, 
    "date"=$6, 
    status=$7, 
    description=$8, 
    payment_method=$9,
    journal_id=$10,
    updated_at=now(),
    updated_by=$11
WHERE id = $1
RETURNING *;

-- name: DeleteTrxExpense :exec
DELETE FROM trx_expenses
WHERE id = $1;
