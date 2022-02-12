package main

import (
	_ "cloudImage/boot"
	_ "cloudImage/router"
	"github.com/gogf/gf/frame/g"
	"os"
	"path/filepath"
)

func main() {
	s := g.Server()

	rootPath, err := os.Getwd()
	if err != nil {
		os.Exit(-1)
	}
	s.SetServerRoot(filepath.Join(rootPath, "public", "static"))
	s.AddStaticPath("/img", filepath.Join(rootPath, "public", "static"))

	s.Run()
}
