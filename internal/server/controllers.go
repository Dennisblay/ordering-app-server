package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	error2 "ordering-server/pkg/error"
	db "ordering-server/pkg/models"
)

func (s *Server) getAllUsersController(ctx *gin.Context) {
	users, err := s.store.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (s *Server) getUserByIdController(ctx *gin.Context) {
	var reqID db.UserRequestByID
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
	email := ctx.Param("email")
	user, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) getUserByPhoneController(ctx *gin.Context) {
	var reqID db.UserRequestByPhone
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	user, err := s.store.GetUserByPhone(ctx, reqID.Phone)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) getUserByEmailAndPasswordController(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.store.GetUserByEmailAndPassword(ctx, req.Email, req.Password)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) createUserController(ctx *gin.Context) {
	var req struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
		Phone        string `json:"phone"`
		Address      string `json:"address"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.store.CreateUser(ctx, req.FirstName, req.LastName, req.Email, req.PasswordHash, req.Phone, req.Address)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserNameController(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.store.UpdateUserName(ctx, id, req.FirstName, req.LastName)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserEmailController(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Email string `json:"email"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.store.UpdateUserEmail(ctx, id, req.Email)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserPhoneController(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Phone string `json:"phone"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.store.UpdateUserPhone(ctx, id, req.Phone)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserAddressController(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Address string `json:"address"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.store.UpdateUserAddress(ctx, id, req.Address)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) updateUserPasswordController(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		PasswordHash string `json:"password_hash"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
		q
	}
	user, err := s.store.UpdateUserPassword(ctx, id, req.PasswordHash)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) deleteUserController(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.store.DeleteUser(ctx, id)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
