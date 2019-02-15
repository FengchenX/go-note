package router

import "agfun/tv/controller"

func initFreeSvc() {
	group := rut.Group("/free-tv")
	group.GET("/", controller.GetFreeTVs)
	group.POST("/", controller.AddFreeTVs)
	group.PUT("/{id}", controller.UpdateFreeTV)
	group.DELETE("/{id}", controller.DelFreeTV)
}
