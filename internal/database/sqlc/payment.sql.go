// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: payment.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payment (order_id, amount, payment_method, payment_date)
VALUES ($1, $2, $3, $4)
    RETURNING id, order_id, amount, payment_method, payment_date, created_at, updated_at
`

type CreatePaymentParams struct {
	OrderID       pgtype.Int4    `json:"order_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	PaymentDate   interface{}    `json:"payment_date"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment,
		arg.OrderID,
		arg.Amount,
		arg.PaymentMethod,
		arg.PaymentDate,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Amount,
		&i.PaymentMethod,
		&i.PaymentDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE
FROM payment
WHERE id = $1
`

func (q *Queries) DeletePayment(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePayment, id)
	return err
}

const getAllPayments = `-- name: GetAllPayments :many
SELECT id, order_id, amount, payment_method, payment_date
FROM payment
`

type GetAllPaymentsRow struct {
	ID            int32          `json:"id"`
	OrderID       pgtype.Int4    `json:"order_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	PaymentDate   interface{}    `json:"payment_date"`
}

func (q *Queries) GetAllPayments(ctx context.Context) ([]GetAllPaymentsRow, error) {
	rows, err := q.db.Query(ctx, getAllPayments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllPaymentsRow{}
	for rows.Next() {
		var i GetAllPaymentsRow
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.Amount,
			&i.PaymentMethod,
			&i.PaymentDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaymentById = `-- name: GetPaymentById :one
SELECT id, order_id, amount, payment_method, payment_date
FROM payment
WHERE id = $1
    LIMIT 1
`

type GetPaymentByIdRow struct {
	ID            int32          `json:"id"`
	OrderID       pgtype.Int4    `json:"order_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	PaymentDate   interface{}    `json:"payment_date"`
}

func (q *Queries) GetPaymentById(ctx context.Context, id int32) (GetPaymentByIdRow, error) {
	row := q.db.QueryRow(ctx, getPaymentById, id)
	var i GetPaymentByIdRow
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Amount,
		&i.PaymentMethod,
		&i.PaymentDate,
	)
	return i, err
}

const getPaymentsByOrderId = `-- name: GetPaymentsByOrderId :many
SELECT id, order_id, amount, payment_method, payment_date
FROM payment
WHERE order_id = $1
`

type GetPaymentsByOrderIdRow struct {
	ID            int32          `json:"id"`
	OrderID       pgtype.Int4    `json:"order_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	PaymentDate   interface{}    `json:"payment_date"`
}

func (q *Queries) GetPaymentsByOrderId(ctx context.Context, orderID pgtype.Int4) ([]GetPaymentsByOrderIdRow, error) {
	rows, err := q.db.Query(ctx, getPaymentsByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPaymentsByOrderIdRow{}
	for rows.Next() {
		var i GetPaymentsByOrderIdRow
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.Amount,
			&i.PaymentMethod,
			&i.PaymentDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePayment = `-- name: UpdatePayment :one
UPDATE payment
SET amount = $2,
    payment_method = $3,
    payment_date = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, order_id, amount, payment_method, payment_date, created_at, updated_at
`

type UpdatePaymentParams struct {
	ID            int32          `json:"id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	PaymentDate   interface{}    `json:"payment_date"`
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, updatePayment,
		arg.ID,
		arg.Amount,
		arg.PaymentMethod,
		arg.PaymentDate,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Amount,
		&i.PaymentMethod,
		&i.PaymentDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
