package router

import "agfun/auth/controller"

func initUser() {
	group := rut.Group("/user")
	group.POST("/register", controller.CreateUser)
	group.GET("/login/:user-name/:pwd", controller.Login)
	group.POST("/vip", controller.AddVip)
	group.GET("/vip", controller.GetVips)
	group.PUT("/vip/:id", controller.UpdateVip)
	group.DELETE("vip/:id", controller.DelVip)
}
