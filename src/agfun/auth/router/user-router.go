package router

import (
	"agfun/auth/controller"
)

func initUser() {
	group := rut.Group("/users/")

	// swagger:route POST /users users createUser
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
	group.POST("/", controller.CreateUser)

	// swagger:route GET /users/{user-name} users login
	//
	// user login
	//
	// user login
	//
	// responses:
	//   default: genericError
	//   200: User
	//   422: validationError
	group.GET("/:user-name", controller.Login)

	// swagger:route POST /users/{user-id}/roles/{role-id} roles addUserRole
	//
	// add user role
	//
	// add user role
	//
	// responses:
	//   200: User
	group.POST("/:user-id/roles/:role-id", controller.AddUserRole)

	group.GET("/vips", controller.GetVips)
	//group.PUT("/vips/:id", controller.UpdateVip)
	//group.DELETE("vips/:id", controller.DelVip)
}
