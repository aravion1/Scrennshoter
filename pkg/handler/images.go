package handler

import (
	"github.com/aravion1/Scrennshoter/structs"
	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	Page_url string `json:"url" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

type ImageResponse struct {
	image []byte
}

var Request ImageRequest

func (h *Handler) getImageByUrl(c *gin.Context) {
	p := structs.Params{Url: Request.Page_url}
	image := h.services.ImageGenerator.GetImage(p)

	// resp := ImageResponse{image: image}
	c.Writer.Write(image)
	// c.JSON(http.StatusOK, image)
}
