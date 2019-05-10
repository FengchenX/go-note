package service

import (
	"github.com/kataras/iris"
	"service"
)

const maxSize = 1 << 30 // 5MB

type Router struct {
	service.Router
	Svr *FileSvc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Party(path string) iris.Party {
	return r.Svr.Party(path)
}

func (r *Router) Handle() {
	root := r.Party("/files")

	upload := root.Party("/upload")
	upload.Post("/videos", iris.LimitRequestBodySize(maxSize), r.Svr.AddVideo)
}
