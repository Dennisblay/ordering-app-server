// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error)
	CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error)
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error)
	CreateRole(ctx context.Context, name string) (Role, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSocialMedia(ctx context.Context, arg CreateSocialMediaParams) (SocialMedia, error)
	CreateTestimonial(ctx context.Context, arg CreateTestimonialParams) (Testimonial, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserRole(ctx context.Context, arg CreateUserRoleParams) (UserRole, error)
	DeleteBooking(ctx context.Context, id int32) error
	DeleteNotification(ctx context.Context, id int32) error
	DeleteOrder(ctx context.Context, id int32) error
	DeleteOrderItem(ctx context.Context, id int32) error
	DeletePayment(ctx context.Context, id int32) error
	DeleteProduct(ctx context.Context, id int32) error
	DeleteReview(ctx context.Context, id int32) error
	DeleteRole(ctx context.Context, id int32) error
	DeleteSession(ctx context.Context, id int32) error
	DeleteSocialMedia(ctx context.Context, id int32) error
	DeleteTestimonial(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserRole(ctx context.Context, arg DeleteUserRoleParams) error
	GetAllBookings(ctx context.Context) ([]GetAllBookingsRow, error)
	GetAllNotifications(ctx context.Context) ([]GetAllNotificationsRow, error)
	GetAllOrderItems(ctx context.Context) ([]GetAllOrderItemsRow, error)
	GetAllOrders(ctx context.Context) ([]GetAllOrdersRow, error)
	GetAllPayments(ctx context.Context) ([]GetAllPaymentsRow, error)
	GetAllProducts(ctx context.Context) ([]GetAllProductsRow, error)
	GetAllReviews(ctx context.Context) ([]GetAllReviewsRow, error)
	GetAllRoles(ctx context.Context) ([]GetAllRolesRow, error)
	GetAllSessions(ctx context.Context) ([]GetAllSessionsRow, error)
	GetAllSocialMedia(ctx context.Context) ([]GetAllSocialMediaRow, error)
	GetAllTestimonials(ctx context.Context) ([]GetAllTestimonialsRow, error)
	GetAllUserRoles(ctx context.Context) ([]UserRole, error)
	GetBookingById(ctx context.Context, id int32) (GetBookingByIdRow, error)
	GetBookingsByUserId(ctx context.Context, userID pgtype.Int4) ([]GetBookingsByUserIdRow, error)
	GetNotificationById(ctx context.Context, id int32) (GetNotificationByIdRow, error)
	GetNotificationsByUserId(ctx context.Context, userID pgtype.Int4) ([]GetNotificationsByUserIdRow, error)
	GetOrderById(ctx context.Context, id int32) (GetOrderByIdRow, error)
	GetOrderItemById(ctx context.Context, id int32) (GetOrderItemByIdRow, error)
	GetOrderItemsByOrderId(ctx context.Context, orderID pgtype.Int4) ([]GetOrderItemsByOrderIdRow, error)
	GetOrdersByUserId(ctx context.Context, userID pgtype.Int4) ([]GetOrdersByUserIdRow, error)
	GetPaymentById(ctx context.Context, id int32) (GetPaymentByIdRow, error)
	GetPaymentsByOrderId(ctx context.Context, orderID pgtype.Int4) ([]GetPaymentsByOrderIdRow, error)
	GetProductById(ctx context.Context, id int32) (GetProductByIdRow, error)
	GetReviewById(ctx context.Context, id int32) (GetReviewByIdRow, error)
	GetReviewsByProductId(ctx context.Context, productID pgtype.Int4) ([]GetReviewsByProductIdRow, error)
	GetRoleById(ctx context.Context, id int32) (GetRoleByIdRow, error)
	GetRoleByName(ctx context.Context, name string) (GetRoleByNameRow, error)
	GetSessionById(ctx context.Context, id int32) (GetSessionByIdRow, error)
	GetSessionByToken(ctx context.Context, token string) (GetSessionByTokenRow, error)
	GetSocialMediaById(ctx context.Context, id int32) (GetSocialMediaByIdRow, error)
	GetSocialMediaByUserId(ctx context.Context, userID pgtype.Int4) ([]GetSocialMediaByUserIdRow, error)
	GetTestimonialById(ctx context.Context, id int32) (GetTestimonialByIdRow, error)
	GetTestimonialsByUserId(ctx context.Context, userID pgtype.Int4) ([]GetTestimonialsByUserIdRow, error)
	GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (GetUserByEmailAndPasswordRow, error)
	GetUserByEmailOrPassword(ctx context.Context, arg GetUserByEmailOrPasswordParams) (GetUserByEmailOrPasswordRow, error)
	GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error)
	GetUserRolesByRoleId(ctx context.Context, roleID int32) ([]UserRole, error)
	GetUserRolesByUserId(ctx context.Context, userID int32) ([]UserRole, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	UpdateBooking(ctx context.Context, arg UpdateBookingParams) (Booking, error)
	UpdateNotification(ctx context.Context, arg UpdateNotificationParams) (Notification, error)
	UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error)
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (Order, error)
	UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error)
	UpdateRoleName(ctx context.Context, arg UpdateRoleNameParams) (Role, error)
	UpdateSocialMedia(ctx context.Context, arg UpdateSocialMediaParams) (SocialMedia, error)
	UpdateTestimonial(ctx context.Context, arg UpdateTestimonialParams) (Testimonial, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	UpdateUserAddress(ctx context.Context, arg UpdateUserAddressParams) (UpdateUserAddressRow, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (UpdateUserEmailRow, error)
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (UpdateUserNameRow, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
	UpdateUserPhone(ctx context.Context, arg UpdateUserPhoneParams) (UpdateUserPhoneRow, error)
}

var _ Querier = (*Queries)(nil)
