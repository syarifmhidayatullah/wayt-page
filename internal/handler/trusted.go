package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type TrustedHandler struct{ svc service.TrustedService }

func NewTrustedHandler(svc service.TrustedService) *TrustedHandler {
	return &TrustedHandler{svc: svc}
}

type trustedRequest struct {
	Name      string `json:"name"       binding:"required"`
	Emoji     string `json:"emoji"`
	Rating    string `json:"rating"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
}

func (h *TrustedHandler) ListPublic(c *gin.Context) {
	list, err := h.svc.ListActive()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *TrustedHandler) List(c *gin.Context) {
	list, err := h.svc.ListAll()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *TrustedHandler) Create(c *gin.Context) {
	var req trustedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	t, err := h.svc.Create(req.Name, req.Emoji, req.Rating, req.SortOrder)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.Created(c, "restoran berhasil ditambahkan", t)
}

func (h *TrustedHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	var req trustedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	t, err := h.svc.Update(uint(id), req.Name, req.Emoji, req.Rating, req.IsActive, req.SortOrder)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "restoran diupdate", t)
}

func (h *TrustedHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "restoran dihapus", nil)
}
