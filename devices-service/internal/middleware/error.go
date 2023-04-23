package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, c.Errors[0].Error())
	}
}
