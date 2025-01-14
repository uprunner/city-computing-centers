// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: payment.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (order_id, amount, status, payment_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING payment_id, order_id, amount, status, payment_date, created_at, updated_at
`

type CreatePaymentParams struct {
	OrderID     int32         `json:"order_id"`
	Amount      string        `json:"amount"`
	Status      PaymentStatus `json:"status"`
	PaymentDate time.Time     `json:"payment_date"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.OrderID,
		arg.Amount,
		arg.Status,
		arg.PaymentDate,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.OrderID,
		&i.Amount,
		&i.Status,
		&i.PaymentDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payments WHERE payment_id = $1
`

func (q *Queries) DeletePayment(ctx context.Context, paymentID int32) error {
	_, err := q.db.ExecContext(ctx, deletePayment, paymentID)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT payment_id, order_id, amount, status, payment_date, created_at, updated_at
FROM payments
WHERE payment_id = $1
LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, paymentID int32) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, paymentID)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.OrderID,
		&i.Amount,
		&i.Status,
		&i.PaymentDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPayments = `-- name: ListPayments :many
SELECT payment_id, order_id, amount, status, payment_date, created_at, updated_at
FROM payments
ORDER BY payment_id
LIMIT $1
    OFFSET $2
`

type ListPaymentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPayments(ctx context.Context, arg ListPaymentsParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listPayments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.PaymentID,
			&i.OrderID,
			&i.Amount,
			&i.Status,
			&i.PaymentDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePayment = `-- name: UpdatePayment :one
UPDATE payments
SET order_id = $1, amount = $2, status = $3, payment_date = $4, updated_at = $5
WHERE payment_id = $6
RETURNING payment_id, order_id, amount, status, payment_date, created_at, updated_at
`

type UpdatePaymentParams struct {
	OrderID     int32         `json:"order_id"`
	Amount      string        `json:"amount"`
	Status      PaymentStatus `json:"status"`
	PaymentDate time.Time     `json:"payment_date"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	PaymentID   int32         `json:"payment_id"`
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, updatePayment,
		arg.OrderID,
		arg.Amount,
		arg.Status,
		arg.PaymentDate,
		arg.UpdatedAt,
		arg.PaymentID,
	)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.OrderID,
		&i.Amount,
		&i.Status,
		&i.PaymentDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
