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

	// swagger:route POST /users/{user-id}/roles/{role-id} users addUserRole
	//
	// add user role
	//
	// add user role
	//
	// responses:
	//   200: UserRole
	group.POST("/:user-id/roles/:role-id", controller.AddUserRole)
	// swagger:route GET /users/{user-id}/roles users getUserRoles
	//
	// get user roles
	//
	// get user roles
	//
	// responses:
	//   200: []UserRole
	group.GET(":user-id/roles", controller.GetUserRoles)
	// swagger:route UPDATE /users/{user-id}/roles/{role-id} users updateUserRole
	//
	// update user role
	//
	// update user role
	//
	// responses:
	//   200: UserRole
	group.PUT("/:user-id/roles/:role-id", controller.UpdateUserRole)
	// swagger:route DELETE /users/{user-id}/roles/{role-id} users delUserRole
	//
	// delete user role
	//
	// delete user role
	//
	// responses:
	//   200: UserRole
	group.DELETE("/:user-id/roles/:role-id", controller.DelUserRole)
}
