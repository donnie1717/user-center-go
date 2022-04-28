package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"username": username})
}
