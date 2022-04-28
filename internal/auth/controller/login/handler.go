package login

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user-center/internal/auth/service"
)

var requestMap = map[string]string{}

type LoginRO struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	ReqId string `form:"req_id"`
}

// 密码登录
func passwordLoginHandler(c *gin.Context)  {
	var ro LoginRO
	if err := c.ShouldBind(&ro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("请求参数 form:%v\n", ro)

	// 后续校验需要放在数据库
	userService := service.GetUserService()

	success, sessionId := userService.PasswordLogin(ro.Username, ro.Password)
	if success {
		c.SetCookie("session_id", sessionId, 10, "/", "", false, false)
		c.Redirect(http.StatusFound, "/index")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"failed": true,
		"error_msg": "用户名密码错误",
	})
}