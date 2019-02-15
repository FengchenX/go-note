package router

import "agfun/top/controller"

func initBannerSvc() {
	group := rutGroup.Group("/banner")
	group.GET("/", controller.GetFreeTVs)
	//group.POST("/", controller.AddFreeTVs)
	//group.PUT("/{id}", controller.UpdateFreeTV)
	//group.DELETE("/{id}", controller.DelFreeTV)
}
