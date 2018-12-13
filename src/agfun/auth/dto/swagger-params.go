package dto

import (
	"agfun/auth/entity"
	"net/http"
)

// user role
// swagger:parameters listVips addUserRole
type UserRoleParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
	// auth session
	// in: header
	Session string `json:"session"`
	// user id
	// in: path
	UserId string `json:"user-id"`
	// role id
	// in: path
	RoleId string `json:"role-id"`

	// body
	// in: body
	Body UserRole `json:"body"`
}

// CreateUserParams contains all the bound params for the create user operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUser
type CreateUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
	/*Created user object
	  In: body
	*/
	Body *entity.User `json:"body"`
}

// login
// swagger:parameters login
type LoginParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	// UserName
	// In: path
	UserName string `json:"user-name"`

	// Pwd
	// In: query
	Pwd string `json:"pwd"`
}

// add role
// swagger:parameters addRole
type AddRoleParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
	// session
	// In: header
	Session string `json:"session"`
	// role
	// In: body
	Role Role `json:"role"`
}