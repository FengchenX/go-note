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

	free := root.Party("/frees")
	free.Post("", r.Svr.AddVideo)
	free.Get("", r.Svr.GetVideos)
	free.Put("/{id}", r.Svr.UpVideo)
	free.Delete("/{id}", r.Svr.DelVideo)

	//pay := root.Party("/pays")
	//pay.Post("")

}
