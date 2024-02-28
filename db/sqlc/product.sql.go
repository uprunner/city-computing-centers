// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (name, description, price, stock_quantity, seller_id, category_id, is_active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING product_id, name, description, price, stock_quantity, seller_id, category_id, is_active, created_at, updated_at
`

type CreateProductParams struct {
	Name          sql.NullString `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         sql.NullString `json:"price"`
	StockQuantity sql.NullInt32  `json:"stock_quantity"`
	SellerID      sql.NullInt32  `json:"seller_id"`
	CategoryID    sql.NullInt32  `json:"category_id"`
	IsActive      sql.NullBool   `json:"is_active"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StockQuantity,
		arg.SellerID,
		arg.CategoryID,
		arg.IsActive,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQuantity,
		&i.SellerID,
		&i.CategoryID,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, productID)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, name, description, price, stock_quantity, seller_id, category_id, is_active, created_at, updated_at
FROM products
WHERE product_id = $1
LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productID int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQuantity,
		&i.SellerID,
		&i.CategoryID,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT product_id, name, description, price, stock_quantity, seller_id, category_id, is_active, created_at, updated_at
FROM products
ORDER BY product_id
LIMIT $1
    OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.StockQuantity,
			&i.SellerID,
			&i.CategoryID,
			&i.IsActive,
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

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $1, description = $2, price = $3, stock_quantity = $4, seller_id = $5, category_id = $6, is_active = $7, updated_at = $8
WHERE product_id = $9
RETURNING product_id, name, description, price, stock_quantity, seller_id, category_id, is_active, created_at, updated_at
`

type UpdateProductParams struct {
	Name          sql.NullString `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         sql.NullString `json:"price"`
	StockQuantity sql.NullInt32  `json:"stock_quantity"`
	SellerID      sql.NullInt32  `json:"seller_id"`
	CategoryID    sql.NullInt32  `json:"category_id"`
	IsActive      sql.NullBool   `json:"is_active"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	ProductID     int32          `json:"product_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StockQuantity,
		arg.SellerID,
		arg.CategoryID,
		arg.IsActive,
		arg.UpdatedAt,
		arg.ProductID,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQuantity,
		&i.SellerID,
		&i.CategoryID,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}