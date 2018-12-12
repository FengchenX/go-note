package dto

import (
	"agfun/auth/entity"
	"net/http"
)

// user role
// swagger:parameters listVips addUserRole
type UserRole struct {
	// a BarSlice has bars which are strings
	//
	// in: body
	entity.UserRole
	Expire string `json:"expire"`
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
	// auth-session
	// In: header
	Session string `json:"session"`
	// role
	// In: body
	Role Role `json:"role"`
}

// role
// swagger:model
type Role struct {
	entity.Role
	Draft interface{} `json:"draft"`
}
