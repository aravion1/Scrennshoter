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

		token, _ := c.GetQuery("token")

		if token != "test123" {
			newErrorResponse(c, http.StatusBadRequest, "Неверный токен")
		}

		c.Next()

		fmt.Println(token)
		fmt.Println(time.Since(t))
	}
}
