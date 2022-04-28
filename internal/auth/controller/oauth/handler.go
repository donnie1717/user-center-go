package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type AuthorizeRO struct {
	AppId string `form:"app_id"  binding:"required"`
	ResponseType string `form:"response_type"  binding:"required"`
	Scope string `form:"scope" binding:"required"`
	RedirectUrl string `form:"redirect_uri" binding:"required"`
}

func authorizeHandler(c *gin.Context)  {
	var ro AuthorizeRO
	if err := c.ShouldBind(&ro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.Cookie("session_id")
	if err != nil {
		reqId := uuid.New().String()
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"reqId": reqId,
		})
		return
	}

	log.Printf("请求参数 form:%v\n", ro)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
