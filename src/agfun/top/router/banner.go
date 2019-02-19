package router

import "agfun/top/controller"

func initBannerSvc() {
	group := rut.Group("/banner")
	group.GET("", controller.GetBanners)
	//group.POST("/", controller.AddFreeTVs)
	//group.PUT("/{id}", controller.UpdateFreeTV)
	//group.DELETE("/{id}", controller.DelFreeTV)
}
