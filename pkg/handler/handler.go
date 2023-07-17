package handler

import (
	"github.com/aravion1/Scrennshoter/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) GetHandlers() *gin.Engine {
	handler := gin.New()

	handler.Use(Auth())

	handler.POST("get-image-by-url", h.getImageByUrl)
	handler.POST("get-element-image-by-url", h.getElementImageByUrl)
	return handler
}
