package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type body struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func OK(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, body{true, message, data})
}

func Created(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, body{true, message, data})
}

func BadRequest(c *gin.Context, message string, err error) {
	msg := message
	if err != nil {
		msg = message + ": " + err.Error()
	}
	c.JSON(http.StatusBadRequest, body{false, msg, nil})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, body{false, message, nil})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, body{false, "unauthorized", nil})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, body{false, "forbidden", nil})
}

func InternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, body{false, err.Error(), nil})
}
