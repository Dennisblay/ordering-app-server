// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Booking struct {
	ID         int32       `json:"id"`
	UserID     pgtype.Int4 `json:"user_id"`
	EventDate  pgtype.Date `json:"event_date"`
	EventType  pgtype.Text `json:"event_type"`
	GuestCount pgtype.Int4 `json:"guest_count"`
	Location   pgtype.Text `json:"location"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type Notification struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	Message   pgtype.Text `json:"message"`
	SentAt    interface{} `json:"sent_at"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Order struct {
	ID           int32          `json:"id"`
	UserID       pgtype.Int4    `json:"user_id"`
	TotalPrice   pgtype.Numeric `json:"total_price"`
	Status       pgtype.Text    `json:"status"`
	DeliveryDate pgtype.Date    `json:"delivery_date"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type OrderItem struct {
	ID        int32          `json:"id"`
	OrderID   pgtype.Int4    `json:"order_id"`
	ProductID pgtype.Int4    `json:"product_id"`
	Quantity  pgtype.Int4    `json:"quantity"`
	Price     pgtype.Numeric `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Payment struct {
	ID            int32          `json:"id"`
	OrderID       pgtype.Int4    `json:"order_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	PaymentDate   interface{}    `json:"payment_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Product struct {
	ID          int32          `json:"id"`
	Name        pgtype.Text    `json:"name"`
	Description pgtype.Text    `json:"description"`
	Category    pgtype.Text    `json:"category"`
	Price       pgtype.Numeric `json:"price"`
	Stock       pgtype.Int4    `json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Review struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	ProductID pgtype.Int4 `json:"product_id"`
	Rating    pgtype.Int4 `json:"rating"`
	Comment   pgtype.Text `json:"comment"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Role struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	Token     string      `json:"token"`
	ExpiresAt interface{} `json:"expires_at"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type SocialMedia struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	Platform  pgtype.Text `json:"platform"`
	Link      pgtype.Text `json:"link"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Testimonial struct {
	ID        int32       `json:"id"`
	UserID    pgtype.Int4 `json:"user_id"`
	Message   pgtype.Text `json:"message"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type User struct {
	ID                  int32       `json:"id"`
	FirstName           string      `json:"first_name"`
	LastName            string      `json:"last_name"`
	Email               string      `json:"email"`
	Phone               string      `json:"phone"`
	Address             string      `json:"address"`
	PasswordHash        string      `json:"password_hash"`
	PasswordUpdatedAt   interface{} `json:"password_updated_at"`
	ResetToken          pgtype.Text `json:"reset_token"`
	ResetTokenExpiresAt interface{} `json:"reset_token_expires_at"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
}

type UserRole struct {
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}
