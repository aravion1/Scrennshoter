package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		fmt.Println(time.Since(t))
	}
}
