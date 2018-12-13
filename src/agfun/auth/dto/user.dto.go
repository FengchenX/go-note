package dto

import (
	"agfun/auth/entity"
)

// user role
// swagger:model
type UserRole struct {
	entity.UserRole
	Expire string `json:"expire"`
}

// role
// swagger:model
type Role struct {
	entity.Role
	Draft interface{} `json:"draft"`
}
