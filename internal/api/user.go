package api

import (
	"errors"
	dto "github.com/Dennisblay/ordering-app-server/internal/database/models"
	errorhandler "github.com/Dennisblay/ordering-app-server/pkg/error"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func (s *Server) getUsersController(ctx *gin.Context) {
	var req dto.UsersRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
	}

	args := dto.GetUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := s.store.GetUsers(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
	return
}
func (s *Server) getUserByIDController(ctx *gin.Context) {
	var req dto.UserRequestByID
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
	}

	user, err := s.store.GetUserById(ctx, req.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *Server) getUserControllerByEmailOrPhone(ctx *gin.Context) {
	var req dto.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
	}

	args := dto.GetUserByEmailOrPasswordParams{
		ID:    req.ID,
		Email: req.Email,
		Phone: req.Phone,
	}
	user, err := s.store.GetUserByEmailOrPassword(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) getUserByEmailAndPasswordController(ctx *gin.Context) {
	var req dto.UserRequestByEmailAndPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserByEmailAndPassword(ctx, dto.GetUserByEmailAndPasswordParams{Email: req.Email, PasswordHash: req.Password})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) createUserController(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	user, err := s.store.CreateUser(ctx, dto.CreateUserParams{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, PasswordHash: req.PasswordHash, Phone: req.Phone, Address: req.Address})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserNameController(ctx *gin.Context) {
	var reqID dto.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	var req dto.UpdateUserNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserName(ctx, dto.UpdateUserNameParams{ID: reqID.ID, FirstName: req.FirstName, LastName: req.LastName})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserEmailController(ctx *gin.Context) {
	var reqID dto.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	var req dto.UpdateUserEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserEmail(ctx, dto.UpdateUserEmailParams{ID: reqID.ID, Email: req.Email})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserPhoneController(ctx *gin.Context) {
	var reqID dto.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	var req dto.UpdateUserPhoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserPhone(ctx, dto.UpdateUserPhoneParams{ID: reqID.ID, Phone: req.Phone})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserAddressController(ctx *gin.Context) {
	var reqID dto.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	var req dto.UpdateUserAddressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserAddress(ctx, dto.UpdateUserAddressParams{ID: reqID.ID, Address: req.Address})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserPasswordController(ctx *gin.Context) {
	var reqID dto.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
	}

	var reqPassword dto.UpdateUserPasswordRequest
	if err := ctx.ShouldBind(&reqPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}
	// first Check if old password == to existing password
	args := dto.UpdateUserPasswordParams{
		ID:           reqID.ID,
		PasswordHash: reqPassword.PasswordNew,
	}
	user, err := s.store.UpdateUserPassword(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorhandler.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) deleteUserController(ctx *gin.Context) {
	var reqID dto.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}

	err := s.store.DeleteUser(ctx, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(err))
		return
	}
	ctx.Status(http.StatusNoContent)
}
