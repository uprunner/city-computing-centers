// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (order_date, buyer_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING order_id, order_date, buyer_id, created_at, updated_at
`

type CreateOrderParams struct {
	OrderDate sql.NullTime  `json:"order_date"`
	BuyerID   sql.NullInt32 `json:"buyer_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.OrderDate,
		arg.BuyerID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderDate,
		&i.BuyerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders WHERE order_id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderID int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, orderID)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, order_date, buyer_id, created_at, updated_at
FROM orders
WHERE order_id = $1
LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderID int32) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderDate,
		&i.BuyerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_id, order_date, buyer_id, created_at, updated_at
FROM orders
ORDER BY order_id
LIMIT $1
    OFFSET $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.OrderDate,
			&i.BuyerID,
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

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders
SET order_date = $1, buyer_id = $2, updated_at = $3
WHERE order_id = $4
RETURNING order_id, order_date, buyer_id, created_at, updated_at
`

type UpdateOrderParams struct {
	OrderDate sql.NullTime  `json:"order_date"`
	BuyerID   sql.NullInt32 `json:"buyer_id"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	OrderID   int32         `json:"order_id"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder,
		arg.OrderDate,
		arg.BuyerID,
		arg.UpdatedAt,
		arg.OrderID,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderDate,
		&i.BuyerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}