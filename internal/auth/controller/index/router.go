package index

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/auth/controller/middleware"
)

func Routers(e *gin.Engine)  {
	e.GET("/index", middleware.AuthMiddleware(), indexHandler)
}
