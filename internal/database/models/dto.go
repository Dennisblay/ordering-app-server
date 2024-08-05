package db

type UserRequestByID struct {
	ID int32 `uri:"id"`
}

type UserRequest struct {
	Email string `form:"email" json:"email,omitempty"`
	Phone string `form:"phone" json:"phone,omitempty"`
}

type UsersRequest struct {
	Limit  int32 `uri:"limit" json:"limit,omitempty"`
	Offset int32 `uri:"offset" json:"offset,omitempty"`
}

type UpdateUserAddressRequest struct {
	Address string `json:"address" binding:"required"`
}

type UpdateUserPasswordRequest struct {
	PasswordOld string `json:"password_old" binding:"required"`
	PasswordNew string `json:"password_new" binding:"required"`
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
