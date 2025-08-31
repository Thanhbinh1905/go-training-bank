-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfersByFrom :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;

-- name: ListTransfersByTo :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;

-- name: ListTransfersAll :many
SELECT * FROM transfers
ORDER BY id DESC
LIMIT $1 OFFSET $2;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
