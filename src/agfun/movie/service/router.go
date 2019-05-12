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
	root.Post("", r.Svr.AddMovie)
	root.Get("", r.Svr.GetMovies)

	movie := root.Party("/{id}")
	movie.Get("", r.Svr.GetMovie)

	videos := movie.Party("/videos")
	videos.Post("", r.Svr.AddMovieVideo)
}
