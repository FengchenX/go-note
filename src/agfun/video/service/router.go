package service

import (
	"github.com/kataras/iris"
	"service"
)

type Router struct {
	service.Router
	Svr *VideoSvc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Party(path string) iris.Party {
	return r.Svr.Party(path)
}

func (r *Router) Handle() {
	root := r.Party("/videos")
	root.Post("", r.Svr.AddVideo)
	root.Get("", r.Svr.GetVideos)
}
