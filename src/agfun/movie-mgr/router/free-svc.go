package router

import "agfun/movie-mgr/controller"

func initFreeSvc() {
	group := rut.Group("/free")
	group.GET("/", controller.GetFreeMovies)
	group.POST("/", controller.AddFreeMovies)
}
