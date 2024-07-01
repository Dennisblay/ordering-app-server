package db

type UserRequestByID struct {
	ID int32 `uri:"id" binding:"required"`
}

type UserRequestByPhone struct {
	Phone string `uri:"phone" binding:"required"`
}

type UserRequestByEmail struct {
	Phone string `uri:"email" binding:"required"`
}
