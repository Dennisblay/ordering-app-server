package db

type UserRequestByID struct {
	ID int32 `uri:"id" binding:"required"`
}

type UserRequestByEmail struct {
	Email string `uri:"email" binding:"required"`
}

type UserRequestByPhone struct {
	Phone string `uri:"phone" binding:"required"`
}
type UpdateUserAddressRequest struct {
	Address string `uri:"address" binding:"required"`
}

type CreateUserRequest struct {
	ID           int32  `json:"id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Address      string `json:"address" binding:"required"`
}

type UserRequestByEmailAndPassword struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserPhoneRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type UpdateUserEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type UpdateUserNameRequest struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
}
