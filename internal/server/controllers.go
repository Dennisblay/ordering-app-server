package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	types "ordering-server/internal/database/models"
	error2 "ordering-server/pkg/error"
)

func (s *Server) getAllUsersController(ctx *gin.Context) {
	users, err := s.store.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": error2.ErrorResponse(err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (s *Server) getUserByIdController(ctx *gin.Context) {
	var reqID types.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserById(ctx, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) getUserByEmailController(ctx *gin.Context) {
	var reqEmail types.UserRequestByEmail
	if err := ctx.ShouldBindUri(&reqEmail); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserByEmail(ctx, reqEmail.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) getUserByPhoneController(ctx *gin.Context) {
	var reqPhone types.UserRequestByPhone
	if err := ctx.ShouldBindUri(&reqPhone); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserByPhone(ctx, reqPhone.Phone)
	if err != nil {
		ctx.JSON(http.StatusNotFound, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) getUserByEmailAndPasswordController(ctx *gin.Context) {
	var req types.UserRequestByEmailAndPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserByEmailAndPassword(ctx, types.GetUserByEmailAndPasswordParams{Email: req.Email, PasswordHash: req.Password})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) createUserController(ctx *gin.Context) {
	var req types.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.CreateUser(ctx, types.CreateUserParams{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, PasswordHash: req.PasswordHash, Phone: req.Phone, Address: req.Address})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserNameController(ctx *gin.Context) {
	var reqID types.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	var req types.UpdateUserNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserName(ctx, types.UpdateUserNameParams{ID: reqID.ID, FirstName: req.FirstName, LastName: req.LastName})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserEmailController(ctx *gin.Context) {
	var reqID types.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	var req types.UpdateUserEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserEmail(ctx, types.UpdateUserEmailParams{ID: reqID.ID, Email: req.Email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserPhoneController(ctx *gin.Context) {
	var reqID types.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	var req types.UpdateUserPhoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserPhone(ctx, types.UpdateUserPhoneParams{ID: reqID.ID, Phone: req.Phone})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserAddressController(ctx *gin.Context) {
	var reqID types.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	var req types.UpdateUserAddressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error2.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserAddress(ctx, types.UpdateUserAddressParams{ID: reqID.ID, Address: req.Address})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserPasswordController(ctx *gin.Context) {

}

func (s *Server) deleteUserController(ctx *gin.Context) {
	var reqID types.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}

	err := s.store.DeleteUser(ctx, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error2.ErrorResponse(err))
		return
	}
	ctx.Status(http.StatusNoContent)
}
