package middleware

import (
	"github.com/gin-gonic/gin"
	"ApiServer/demo09/handler"
	"ApiServer/demo09/pkg/errno"
	"ApiServer/demo09/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
