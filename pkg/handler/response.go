package handler

import (
	b64 "encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:""`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}

func response(c *gin.Context, image []byte, err error, request ImageRequest) {
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if request.IsRow {
		c.JSON(http.StatusOK, b64.StdEncoding.EncodeToString(image))
		return
	}

	c.Writer.Write(image)
	return
}
