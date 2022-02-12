package router

import (
	"cloudImage/app/api"
	"cloudImage/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Use(service.Middleware.Auth)
	s.Group("/img", func(group *ghttp.RouterGroup) {
		group.ALL("/upload", api.Image.Upload)
	})
}
