package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getImageByUrl(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
