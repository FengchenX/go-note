package router

import "agfun/auth/controller"

func initResources() {
	group := rut.Group("/resources")

	// swagger:route GET /resources resources getResources
	//
	// get resources
	//
	// get resources
	//
	// responses:
	//   200: null
	group.GET("", controller.GetResources)
}
