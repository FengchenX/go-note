package router

import "agfun/auth/controller"

func initUser() {
	group := rut.Group("/user")
	group.POST("/register", controller.CreateUser)
}
