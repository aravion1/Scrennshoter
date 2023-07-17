package handler

import (
	b64 "encoding/base64"
	"net/http"

	"github.com/aravion1/Scrennshoter/structs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	IsFull:   true,
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

	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if Request.IsRow {
		c.JSON(http.StatusOK, b64.StdEncoding.EncodeToString(image))
		return
	}

	c.Writer.Write(image)
	return
}
