package router

import (
	"agfun/auth/controller"
)

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
	// swagger:route POST /users/roles roles addRole
	//
	// add role
	//
	// add role
	//
	// responses:
	//   200: Role
	group.POST("/roles", controller.AddRole)
	// swagger:route POST /users/roles roles delRole
	//
	// delete role
	//
	// delete role
	//
	// responses:
	//   200: nil
	group.DELETE("/roles/:id", controller.DelRole)
	group.PUT("/roles/:id", controller.UpdateRole)
	group.GET("roles", controller.GetRoles)

	// swagger:route POST /users/{id}/roles/{id} roles addUserRole
	//
	// delete role
	//
	// delete role
	//
	// responses:
	//   200: nil
	group.POST("/vips", controller.AddUserRole)
	group.GET("/vips", controller.GetVips)
	group.PUT("/vips/:id", controller.UpdateVip)
	group.DELETE("vips/:id", controller.DelVip)
}
