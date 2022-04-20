package main

import (
	"user-center/app/auth/controller/oauth"
	"user-center/routers"
)

func main() {
	routers.Include(oauth.Routers)

	r := routers.Init()
	r.LoadHTMLGlob("templates/*")

	r.Run()
}


