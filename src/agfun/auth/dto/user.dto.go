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

// createUserResp
// swagger:response createUserResp
type CreateUserResp struct {

}
