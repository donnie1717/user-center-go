package main

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/auth/controller/index"
	"user-center/internal/auth/controller/login"
	"user-center/internal/auth/controller/oauth"
	routers2 "user-center/pkg/routers"
)

func main() {
	r := gin.Default()

	// 路由注册
	routers2.Include(oauth.Routers, login.Routers, index.Routers)
	routers2.Init(r)

	// 模板设置
	r.LoadHTMLGlob("web/auth/templates/*")



	// ORM初始化

	r.Run()
}
