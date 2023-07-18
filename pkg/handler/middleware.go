package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		_, err := getBearer(c)

		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Next()

		fmt.Println(time.Since(t))
	}
}

func getBearer(c *gin.Context) (string, error) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		return "", errors.New("authentication failed")
	}

	bearer_slice := strings.Split(bearer, " ")
	if len(bearer_slice) < 2 {
		return "", errors.New("authentication failed")
	}

	return bearer_slice[1], nil
}
