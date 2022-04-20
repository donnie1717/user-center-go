package oauth

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine)  {
	e.GET("/oauth/authorize", authorizeHandler)
	e.POST("/oauth/login", passwordLoginHandler)
}