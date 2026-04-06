package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type AuthHandler struct{ svc service.AuthService }

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	token, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "login berhasil", gin.H{"token": token})
}
