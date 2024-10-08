// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: review.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createReview = `-- name: CreateReview :one
INSERT INTO review (user_id, product_id, rating, comment)
VALUES ($1, $2, $3, $4)
    RETURNING id, user_id, product_id, rating, comment, created_at, updated_at
`

type CreateReviewParams struct {
	UserID    pgtype.Int4 `json:"user_id"`
	ProductID pgtype.Int4 `json:"product_id"`
	Rating    pgtype.Int4 `json:"rating"`
	Comment   pgtype.Text `json:"comment"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, createReview,
		arg.UserID,
		arg.ProductID,
		arg.Rating,
		arg.Comment,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE
FROM review
WHERE id = $1
`

func (q *Queries) DeleteReview(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteReview, id)
	return err
}

const getAllReviews = `-- name: GetAllReviews :many
SELECT id, user_id, product_id, rating, comment
FROM review
`

type GetAllReviewsRow struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	ProductID pgtype.Int4 `json:"product_id"`
	Rating    pgtype.Int4 `json:"rating"`
	Comment   pgtype.Text `json:"comment"`
}

func (q *Queries) GetAllReviews(ctx context.Context) ([]GetAllReviewsRow, error) {
	rows, err := q.db.Query(ctx, getAllReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllReviewsRow{}
	for rows.Next() {
		var i GetAllReviewsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.Rating,
			&i.Comment,
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

const getReviewById = `-- name: GetReviewById :one
SELECT id, user_id, product_id, rating, comment
FROM review
WHERE id = $1
    LIMIT 1
`

type GetReviewByIdRow struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	ProductID pgtype.Int4 `json:"product_id"`
	Rating    pgtype.Int4 `json:"rating"`
	Comment   pgtype.Text `json:"comment"`
}

func (q *Queries) GetReviewById(ctx context.Context, id int32) (GetReviewByIdRow, error) {
	row := q.db.QueryRow(ctx, getReviewById, id)
	var i GetReviewByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
	)
	return i, err
}

const getReviewsByProductId = `-- name: GetReviewsByProductId :many
SELECT id, user_id, product_id, rating, comment
FROM review
WHERE product_id = $1
`

type GetReviewsByProductIdRow struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	ProductID pgtype.Int4 `json:"product_id"`
	Rating    pgtype.Int4 `json:"rating"`
	Comment   pgtype.Text `json:"comment"`
}

func (q *Queries) GetReviewsByProductId(ctx context.Context, productID pgtype.Int4) ([]GetReviewsByProductIdRow, error) {
	rows, err := q.db.Query(ctx, getReviewsByProductId, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetReviewsByProductIdRow{}
	for rows.Next() {
		var i GetReviewsByProductIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.Rating,
			&i.Comment,
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

const updateReview = `-- name: UpdateReview :one
UPDATE review
SET rating = $2,
    comment = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, product_id, rating, comment, created_at, updated_at
`

type UpdateReviewParams struct {
	ID      int32       `json:"id"`
	Rating  pgtype.Int4 `json:"rating"`
	Comment pgtype.Text `json:"comment"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, updateReview, arg.ID, arg.Rating, arg.Comment)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
