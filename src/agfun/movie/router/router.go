package router

import (
	"github.com/kataras/iris"
	"router"
)

type Router struct {
	router.Router
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Party(path string) iris.Party {
	return r.App.Party(path)
}

func (r *Router) Handle() {
	root := r.Party("/movies")
	free := root.Party("/frees")
	free.Post("", )
}
