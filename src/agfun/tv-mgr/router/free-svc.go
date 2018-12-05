package router

import "agfun/tv-mgr/controller"

func initFreeSvc() {
	group := rut.Group("/free-movie")
	group.GET("/", controller.GetFreeTV)
	group.POST("/", controller.AddFreeTV)
	group.PUT("/{id}", controller.UpdateFreeTV)
	group.DELETE("/{id}", controller.DelFreeTV)
}
