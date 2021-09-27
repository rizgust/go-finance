INSERT INTO m_book_periods (
    onwer_id, 
    code, 
    "name", 
    description, 
    status, 
    start_date, 
    end_date
    created_by,
    updated_by
) 
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetInvoice :one
SELECT * FROM m_book_periods
WHERE id = $1 LIMIT 1;

-- name: GetInvoiceForUpdate :one
SELECT * FROM m_book_periods
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListInvoices :many
SELECT * FROM m_book_periods
WHERE owner_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateInvoice :one
UPDATE m_book_periods
SET owner_id=$2, 
    code=$3, 
    "name"=$4, 
    description=$5, 
    status=$6, 
    start_date=$7, 
    end_date=$8, 
    updated_at=now(),
    updated_by=$9
WHERE id = $1
RETURNING *;

-- name: DeleteInvoice :exec
DELETE FROM m_book_periods
WHERE id = $1;
