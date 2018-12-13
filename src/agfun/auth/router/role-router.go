package router

import "agfun/auth/controller"

func initRoles() {
	roles := rut.Group("/roles")
	// swagger:route POST /users/roles roles addRole
	//
	// add role
	//
	// add role
	//
	// responses:
	//   200: Role
	roles.POST("/", controller.AddRole)
	// swagger:route POST /users/roles roles delRole
	//
	// delete role
	//
	// delete role
	roles.DELETE("/:id", controller.DelRole)
	roles.PUT("/:id", controller.UpdateRole)
	roles.GET("/", controller.GetRoles)
}
