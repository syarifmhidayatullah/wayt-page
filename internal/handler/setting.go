package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/response"
)

type SettingHandler struct{ svc service.SettingService }

func NewSettingHandler(svc service.SettingService) *SettingHandler {
	return &SettingHandler{svc: svc}
}

func (h *SettingHandler) GetPublic(c *gin.Context) {
	m, err := h.svc.GetMap()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", m)
}

func (h *SettingHandler) List(c *gin.Context) {
	list, err := h.svc.ListAll()
	if err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "success", list)
}

func (h *SettingHandler) Update(c *gin.Context) {
	var req struct {
		Key   string `json:"key"   binding:"required"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request", err)
		return
	}
	if err := h.svc.Set(req.Key, req.Value); err != nil {
		response.InternalError(c, err)
		return
	}
	response.OK(c, "setting disimpan", nil)
}
