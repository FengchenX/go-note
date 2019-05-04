package service

import (
	"github.com/kataras/iris"
	"service"
)

type Router struct {
	service.Router
	Svr *MovieSvc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Party(path string) iris.Party {
	return r.Svr.Party(path)
}

func (r *Router) Handle() {
	root := r.Party("/movies")

	free := root.Party("/frees")
	free.Post("", r.Svr.AddFree)
	free.Get("", r.Svr.GetFrees)
	free.Put("/{id}", r.Svr.UpdateFree)
	free.Delete("/{id}", r.Svr.DeleteFree)

	//pay := root.Party("/pays")
	//pay.Post("")

}
