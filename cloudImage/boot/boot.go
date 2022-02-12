package boot

import (
	"cloudImage/app/service"
	_ "cloudImage/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	s.Use(service.Middleware.CORS)
}
