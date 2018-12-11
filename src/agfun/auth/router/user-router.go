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
	//   200: User
	//   422: validationError
	group.POST("/register", controller.CreateUser)

	// swagger:route GET /users/login/{user-name} users login
	//
	// user login
	//
	// user login
	//
	// responses:
	//   default: genericError
	//   200: User
	//   422: validationError
	group.GET("/login/:user-name", controller.Login)
	group.POST("/vips", controller.AddVip)
	group.GET("/vips", controller.GetVips)
	group.PUT("/vips/:id", controller.UpdateVip)
	group.DELETE("vips/:id", controller.DelVip)
}
