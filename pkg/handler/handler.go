package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetHandlers() *gin.Engine {
	handler := gin.New()
	handler.Use(Auth())
	handler.GET("get-image-by-url", h.getImageByUrl)
	return handler
}
