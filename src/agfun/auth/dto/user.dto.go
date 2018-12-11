package dto

import (
	"agfun/auth/entity"
	"net/http"
)

// swagger:parameters listVips addBars
type VipUser struct {
	// a BarSlice has bars which are strings
	//
	// in: body
	entity.VipUser
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
