INSERT INTO invoices (
    onwer_id, 
    user_id, 
    "number", 
    ar_id, 
    status, 
    amount, 
    amount_paid, 
    "date", 
    due_date, 
    additional_info, 
    payment_id, 
    discount_id, 
    created_at,
    period_id, 
    class_program_id, 
    class_level_id, 
    class_specialization_id, 
    male, 
    recurring_type, 
    recurring_period, 
    installment, 
    mutation, 
    boarding, 
    admission_line_id, 
    admission_batch_id,
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28
) RETURNING *;

-- name: GetInvoice :one
SELECT * FROM invoices
WHERE id = $1 LIMIT 1;

-- name: GetInvoiceForUpdate :one
SELECT * FROM invoices
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListInvoices :many
SELECT * FROM invoices
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateInvoice :one
UPDATE invoices
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    alias=$5,
    updated_at=now(),
    updated_by=$6
WHERE id = $1
RETURNING *;

-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1;
