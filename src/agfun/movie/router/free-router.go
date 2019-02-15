package router

import "agfun/movie/controller"

func initFreeSvc() {
	group := rut.Group("/free-movies")

	// swagger:route GET /free-movies movies getFreeMovies
	//
	// get free movies
	//
	// get free movies
	//
	// responses:
	//   200: FreeMovie
	group.GET("", controller.GetFreeMovies)

	// swagger:route POST /free-movies movies addFreeMovies
	//
	// add free movies addFreeMovies
	//
	// add free movies
	//
	// responses:
	//   200: FreeMovie
	group.POST("", controller.AddFreeMovies)
	group.PUT("/{id}", controller.UpdateFreeMovie)
	group.DELETE("/{id}", controller.DelFreeMovie)
}
