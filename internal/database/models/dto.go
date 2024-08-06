package db

type UserRequestByID struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type UserRequest struct {
	ID    int32  `form:"id" json:"id,omitempty" binding:"min=1"`
	Email string `form:"email" json:"email,omitempty"`
	Phone string `form:"phone" json:"phone,omitempty"`
}

type UsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type UpdateUserAddressRequest struct {
	Address string `json:"address" binding:"required"`
}

type UpdateUserPasswordRequest struct {
	PasswordOld string `json:"password_old" binding:"required"`
	PasswordNew string `json:"password_new" binding:"required"`
}

type CreateUserRequest struct {
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
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
