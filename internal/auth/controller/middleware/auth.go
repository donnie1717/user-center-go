package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-center/internal/auth/service"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("session_id")

		if err != nil {
			c.HTML(http.StatusFound, "login.tmpl", gin.H{"failed": true, "error_msg": "当前无登录状态，请先完成登录！"})
			c.Abort()
			return
		}

		username := service.GetUserService().GetBySessionId(sessionId)
		if username == "" {
			c.HTML(http.StatusFound, "login.tmpl", gin.H{"failed": true, "error_msg": "当前无登录状态，请先完成登录！"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}