package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 错误处理
		defer func() {
			for _, err := range c.Errors {
				c.AbortWithStatusJSON(c.Writer.Status(), gin.H{
					"code":    c.Writer.Status(),
					"message": http.StatusText(c.Writer.Status()),
					"cause":   err.Error(),
				})
				return
			}
		}()
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}
