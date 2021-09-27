INSERT INTO invoices (
    owner_id, 
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
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27
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
    user_id=$3, 
    "number"=$4, 
    ar_id=$5, 
    status=$6, 
    amount=$7, 
    amount_paid=$8, 
    "date"=$9, 
    due_date=$10, 
    additional_info=$11, 
    payment_id=$12, 
    discount_id=$13,
    period_id=$14, 
    class_program_id=$15, 
    class_level_id=$16, 
    class_specialization_id=$18, 
    male=$19, 
    recurring_type=$20, 
    recurring_period=$21, 
    installment=$22, 
    mutation=$23, 
    boarding=$24, 
    admission_line_id=$25, 
    admission_batch_id=$26,
    updated_at=now(),
    updated_by=$27
WHERE id = $1
RETURNING *;

-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1;
