package index

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine)  {
	e.GET("/index", indexHandler)
}
