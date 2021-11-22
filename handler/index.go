package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) indexPage(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": "hello, world!",
	})
}
