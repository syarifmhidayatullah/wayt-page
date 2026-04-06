package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type PricingHandler struct{ svc service.PricingService }

func NewPricingHandler(svc service.PricingService) *PricingHandler {
	return &PricingHandler{svc: svc}
}

type pricingRequest struct {
	Name         string   `json:"name"          binding:"required"`
	PriceMonthly int      `json:"price_monthly"`
	PriceYearly  int      `json:"price_yearly"`
	Features     []string `json:"features"`
	IsPopular    bool     `json:"is_popular"`
	IsActive     bool     `json:"is_active"`
	SortOrder    int      `json:"sort_order"`
}

// Public
func (h *PricingHandler) ListPublic(c *gin.Context) {
	list, err := h.svc.ListActive()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

// Internal
func (h *PricingHandler) List(c *gin.Context) {
	list, err := h.svc.ListAll()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *PricingHandler) Create(c *gin.Context) {
	var req pricingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	p, err := h.svc.Create(req.Name, req.PriceMonthly, req.PriceYearly, req.Features, req.IsPopular, req.SortOrder)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.Created(c, "paket berhasil dibuat", p)
}

func (h *PricingHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	var req pricingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	p, err := h.svc.Update(uint(id), req.Name, req.PriceMonthly, req.PriceYearly, req.Features, req.IsPopular, req.IsActive, req.SortOrder)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "paket berhasil diupdate", p)
}

func (h *PricingHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "paket dihapus", nil)
}
