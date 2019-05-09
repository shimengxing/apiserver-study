package middleware

import (
	"apiserver-study/handler"
	"apiserver-study/pkg/errno"
	"apiserver-study/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//解析jwt
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			//防止调用挂起的处理程序，确保程序可以处理其他程序
			c.Abort()
			return
		}

		c.Next()
	}
}
