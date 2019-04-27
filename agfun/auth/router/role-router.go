package router

import "agfun/auth/controller"

func initRoles() {
	roles := rut.Group("/roles")
	// swagger:route POST /roles roles addRole
	//
	// add role
	//
	// add role
	//
	// responses:
	//   200: Role
	roles.POST("", controller.AddRole)
	// swagger:route POST /roles roles delRole
	//
	// delete role
	//
	// delete role
	roles.DELETE("/:id", controller.DelRole)
	roles.PUT("/:id", controller.UpdateRole)
	// swagger:route GET /roles roles getRoles
	//
	// get roles
	//
	// get roles
	//
	// responses:
	//   200: []Role
	roles.GET("", controller.GetRoles)
}
