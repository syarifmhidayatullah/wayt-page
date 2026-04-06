package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type AuthHandler struct{ svc service.AuthService }

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) ListAdmins(c *gin.Context) {
	list, err := h.svc.ListAdmins()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *AuthHandler) CreateAdmin(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	u, err := h.svc.CreateAdmin(req.Username, req.Password)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.Created(c, "admin berhasil dibuat", u)
}

func (h *AuthHandler) UpdatePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	var req struct {
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	if err := h.svc.UpdatePassword(uint(id), req.Password); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "password diupdate", nil)
}

func (h *AuthHandler) DeleteAdmin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	if err := h.svc.DeleteAdmin(uint(id)); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "admin dihapus", nil)
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
