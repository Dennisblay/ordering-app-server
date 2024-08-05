// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: social_media.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSocialMedia = `-- name: CreateSocialMedia :one
INSERT INTO social_media (user_id, platform, link)
VALUES ($1, $2, $3)
    RETURNING id, user_id, platform, link, created_at, updated_at
`

type CreateSocialMediaParams struct {
	UserID   pgtype.Int4 `json:"user_id"`
	Platform pgtype.Text `json:"platform"`
	Link     pgtype.Text `json:"link"`
}

func (q *Queries) CreateSocialMedia(ctx context.Context, arg CreateSocialMediaParams) (SocialMedia, error) {
	row := q.db.QueryRow(ctx, createSocialMedia, arg.UserID, arg.Platform, arg.Link)
	var i SocialMedia
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSocialMedia = `-- name: DeleteSocialMedia :exec
DELETE
FROM social_media
WHERE id = $1
`

func (q *Queries) DeleteSocialMedia(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteSocialMedia, id)
	return err
}

const getAllSocialMedia = `-- name: GetAllSocialMedia :many
SELECT id, user_id, platform, link
FROM social_media
`

type GetAllSocialMediaRow struct {
	ID       int32       `json:"id"`
	UserID   pgtype.Int4 `json:"user_id"`
	Platform pgtype.Text `json:"platform"`
	Link     pgtype.Text `json:"link"`
}

func (q *Queries) GetAllSocialMedia(ctx context.Context) ([]GetAllSocialMediaRow, error) {
	rows, err := q.db.Query(ctx, getAllSocialMedia)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllSocialMediaRow{}
	for rows.Next() {
		var i GetAllSocialMediaRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Platform,
			&i.Link,
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

const getSocialMediaById = `-- name: GetSocialMediaById :one
SELECT id, user_id, platform, link
FROM social_media
WHERE id = $1
    LIMIT 1
`

type GetSocialMediaByIdRow struct {
	ID       int32       `json:"id"`
	UserID   pgtype.Int4 `json:"user_id"`
	Platform pgtype.Text `json:"platform"`
	Link     pgtype.Text `json:"link"`
}

func (q *Queries) GetSocialMediaById(ctx context.Context, id int32) (GetSocialMediaByIdRow, error) {
	row := q.db.QueryRow(ctx, getSocialMediaById, id)
	var i GetSocialMediaByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Link,
	)
	return i, err
}

const getSocialMediaByUserId = `-- name: GetSocialMediaByUserId :many
SELECT id, user_id, platform, link
FROM social_media
WHERE user_id = $1
`

type GetSocialMediaByUserIdRow struct {
	ID       int32       `json:"id"`
	UserID   pgtype.Int4 `json:"user_id"`
	Platform pgtype.Text `json:"platform"`
	Link     pgtype.Text `json:"link"`
}

func (q *Queries) GetSocialMediaByUserId(ctx context.Context, userID pgtype.Int4) ([]GetSocialMediaByUserIdRow, error) {
	rows, err := q.db.Query(ctx, getSocialMediaByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSocialMediaByUserIdRow{}
	for rows.Next() {
		var i GetSocialMediaByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Platform,
			&i.Link,
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

const updateSocialMedia = `-- name: UpdateSocialMedia :one
UPDATE social_media
SET platform = $2,
    link = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, platform, link, created_at, updated_at
`

type UpdateSocialMediaParams struct {
	ID       int32       `json:"id"`
	Platform pgtype.Text `json:"platform"`
	Link     pgtype.Text `json:"link"`
}

func (q *Queries) UpdateSocialMedia(ctx context.Context, arg UpdateSocialMediaParams) (SocialMedia, error) {
	row := q.db.QueryRow(ctx, updateSocialMedia, arg.ID, arg.Platform, arg.Link)
	var i SocialMedia
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
