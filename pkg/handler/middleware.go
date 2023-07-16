package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		if err := c.BindJSON(&Request); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if Request.Token != "test123" {
			newErrorResponse(c, http.StatusBadRequest, "Неверный токен")
			return
		}

		c.Next()

		fmt.Println(time.Since(t))
	}
}
