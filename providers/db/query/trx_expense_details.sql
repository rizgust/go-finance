INSERT INTO trx_expense_details (
    trx_expenses_id, 
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

-- name: GetTrxExpenseDetail :one
SELECT * FROM trx_expense_details
WHERE id = $1 LIMIT 1;

-- name: GetTrxExpenseDetailForUpdate :one
SELECT * FROM trx_expense_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrxExpenseDetails :many
SELECT * FROM trx_expense_details
WHERE trx_expenses_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTrxExpenseDetail :one
UPDATE trx_expense_details
SET trx_expenses_id=$2, 
    amount=$3, 
    is_source=$4, 
    account_id=$5, 
    model=$6, 
    model_id=$7,
    updated_at=now(),
    updated_by=$8
WHERE id = $1
RETURNING *;

-- name: DeleteTrxExpenseDetail :exec
DELETE FROM trx_expense_details
WHERE id = $1;
