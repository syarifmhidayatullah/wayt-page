package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type LeadHandler struct{ svc service.LeadService }

func NewLeadHandler(svc service.LeadService) *LeadHandler {
	return &LeadHandler{svc: svc}
}

func (h *LeadHandler) Submit(c *gin.Context) {
	var req struct {
		FullName       string `json:"full_name"       binding:"required"`
		RestaurantName string `json:"restaurant_name" binding:"required"`
		Email          string `json:"email"           binding:"required"`
		Phone          string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	l, err := h.svc.Submit(req.FullName, req.RestaurantName, req.Email, req.Phone)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.Created(c, "Terima kasih! Kami akan menghubungi Anda segera.", l)
}

func (h *LeadHandler) List(c *gin.Context) {
	list, err := h.svc.ListAll()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *LeadHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "lead dihapus", nil)
}
