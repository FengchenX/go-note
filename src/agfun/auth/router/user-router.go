package router

import "agfun/auth/controller"

func initUser() {
	group := rut.Group("/users")

	// swagger:route POST /users/register users createUser
	//
	// register user account
	//
	// this will register user account
	//
	//
	// responses:
	//   default: genericError
	//   200: createUserResp
	//   422: validationError
	group.POST("/register", controller.CreateUser)
	group.GET("/login/:user-name/:pwd", controller.Login)
	group.POST("/vips", controller.AddVip)
	group.GET("/vips", controller.GetVips)
	group.PUT("/vips/:id", controller.UpdateVip)
	group.DELETE("vips/:id", controller.DelVip)
}
