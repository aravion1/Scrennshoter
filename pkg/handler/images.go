package handler

import (
	"github.com/aravion1/Scrennshoter/structs"
	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	Page_url string `json:"url" binding:"required"`
	IsFull   bool   `json:"is_full"`
	Token    string `json:"token" binding:"required"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	IsRow    bool   `json:"is_row"`
	Viewport string `json:"viewport"`
	Selector string `json:"selector"`
}

type ImageResponse struct {
	image []byte
}

var Request = ImageRequest{
	IsFull:   false,
	Width:    1920,
	Height:   1080,
	IsRow:    false,
	Selector: "body",
}

func (h *Handler) getImageByUrl(c *gin.Context) {
	p := structs.Params{
		Url:    Request.Page_url,
		IsFull: Request.IsFull,
		Width:  Request.Width,
		Height: Request.Height,
		Sel:    Request.Selector,
	}

	image, err := h.services.ImageGenerator.GetImage(p)

	response(c, image, err)
	return
}

func (h *Handler) getElementImageByUrl(c *gin.Context) {
	p := structs.Params{
		Url:    Request.Page_url,
		IsFull: Request.IsFull,
		Width:  Request.Width,
		Height: Request.Height,
		Sel:    Request.Selector,
	}

	image, err := h.services.ImageGenerator.GetElementImageByUrl(p)

	response(c, image, err)
	return
}
