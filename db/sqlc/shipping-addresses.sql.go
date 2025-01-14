// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: shipping-addresses.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createShippingAddress = `-- name: CreateShippingAddress :one
INSERT INTO shipping_addresses (user_id, address, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING address_id, user_id, address, created_at, updated_at
`

type CreateShippingAddressParams struct {
	UserID    int32        `json:"user_id"`
	Address   string       `json:"address"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateShippingAddress(ctx context.Context, arg CreateShippingAddressParams) (ShippingAddress, error) {
	row := q.db.QueryRowContext(ctx, createShippingAddress,
		arg.UserID,
		arg.Address,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i ShippingAddress
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteShippingAddress = `-- name: DeleteShippingAddress :exec
DELETE FROM shipping_addresses WHERE address_id = $1
`

func (q *Queries) DeleteShippingAddress(ctx context.Context, addressID int32) error {
	_, err := q.db.ExecContext(ctx, deleteShippingAddress, addressID)
	return err
}

const getShippingAddress = `-- name: GetShippingAddress :one
SELECT address_id, user_id, address, created_at, updated_at
FROM shipping_addresses
WHERE address_id = $1
LIMIT 1
`

func (q *Queries) GetShippingAddress(ctx context.Context, addressID int32) (ShippingAddress, error) {
	row := q.db.QueryRowContext(ctx, getShippingAddress, addressID)
	var i ShippingAddress
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listShippingAddresses = `-- name: ListShippingAddresses :many
SELECT address_id, user_id, address, created_at, updated_at
FROM shipping_addresses
ORDER BY address_id
LIMIT $1
    OFFSET $2
`

type ListShippingAddressesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListShippingAddresses(ctx context.Context, arg ListShippingAddressesParams) ([]ShippingAddress, error) {
	rows, err := q.db.QueryContext(ctx, listShippingAddresses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShippingAddress
	for rows.Next() {
		var i ShippingAddress
		if err := rows.Scan(
			&i.AddressID,
			&i.UserID,
			&i.Address,
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

const updateShippingAddress = `-- name: UpdateShippingAddress :one
UPDATE shipping_addresses
SET user_id = $1, address = $2, updated_at = $3
WHERE address_id = $4
RETURNING address_id, user_id, address, created_at, updated_at
`

type UpdateShippingAddressParams struct {
	UserID    int32        `json:"user_id"`
	Address   string       `json:"address"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	AddressID int32        `json:"address_id"`
}

func (q *Queries) UpdateShippingAddress(ctx context.Context, arg UpdateShippingAddressParams) (ShippingAddress, error) {
	row := q.db.QueryRowContext(ctx, updateShippingAddress,
		arg.UserID,
		arg.Address,
		arg.UpdatedAt,
		arg.AddressID,
	)
	var i ShippingAddress
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
