package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-center/internal/auth/service"
)

func indexHandler(ctx *gin.Context) {

	sessionId, err := ctx.Cookie("session_id")

	if err != nil {
		ctx.HTML(http.StatusFound, "login.tmpl", gin.H{"failed": true, "error_msg": "当前无登录状态，请先完成登录！"})
		return
	}

	username := service.GetUserService().GetBySessionId(sessionId)
	if username == "" {
		ctx.HTML(http.StatusFound, "login.tmpl", gin.H{"failed": true, "error_msg": "当前无登录状态，请先完成登录！"})
		return
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"username": username})
}
