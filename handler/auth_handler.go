package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotoism/football-go/dto"
	"github.com/pranotoism/football-go/service"
	"github.com/pranotoism/football-go/util"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authService.Register(req)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "User registered successfully", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.authService.Login(req)
	if err != nil {
		util.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Login successful", dto.AuthResponse{Token: token})
}
