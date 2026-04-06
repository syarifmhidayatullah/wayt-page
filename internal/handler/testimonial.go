package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type TestimonialHandler struct{ svc service.TestimonialService }

func NewTestimonialHandler(svc service.TestimonialService) *TestimonialHandler {
	return &TestimonialHandler{svc: svc}
}

type testimonialRequest struct {
	Name       string `json:"name"       binding:"required"`
	Restaurant string `json:"restaurant"`
	Quote      string `json:"quote"      binding:"required"`
	Phone      string `json:"phone"`
	Rating     int    `json:"rating"`
	SortOrder  int    `json:"sort_order"`
	IsActive   bool   `json:"is_active"`
}

func (h *TestimonialHandler) ListPublic(c *gin.Context) {
	list, err := h.svc.ListActive()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *TestimonialHandler) List(c *gin.Context) {
	list, err := h.svc.ListAll()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *TestimonialHandler) Create(c *gin.Context) {
	var req testimonialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	rating := req.Rating
	if rating <= 0 {
		rating = 5
	}
	t, err := h.svc.Create(req.Name, req.Restaurant, req.Quote, req.Phone, rating, req.SortOrder)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.Created(c, "testimoni berhasil ditambahkan", t)
}

func (h *TestimonialHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	var req testimonialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	t, err := h.svc.Update(uint(id), req.Name, req.Restaurant, req.Quote, req.Phone, req.Rating, req.SortOrder, req.IsActive)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "testimoni diupdate", t)
}

func (h *TestimonialHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id", err)
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}
	response.OK(c, "testimoni dihapus", nil)
}
