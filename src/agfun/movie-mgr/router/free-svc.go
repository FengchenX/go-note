package router

import "agfun/movie-mgr/controller"

func initFreeSvc() {
	group := rut.Group("/free-movie")
	group.GET("/", controller.GetFreeMovies)
	group.POST("/", controller.AddFreeMovies)
	group.PUT("/{id}", controller.UpdateFreeMovie)
	group.DELETE("/{id}", controller.DelFreeMovie)
}
