package entity

import "time"

// Role
// swagger:model
type Role struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Label      string    `json:"label"`
	Desc       string    `json:"desc"`
	CreateTime time.Time `json:"create_time"`
}

// user role
// swagger:model
type UserRole struct {
	ID     string    `json:"id"`
	UserID string    `json:"user_id"`
	RoleID string    `json:"role_id"`
	Level  int       `json:"level"`
	Expire time.Time `json:"expire"`
}

// group role
// swagger:model
type GroupRole struct {
	ID      string    `json:"id"`
	GroupID string    `json:"group_id"`
	UserID  string    `json:"user_id"`
	RoleID  string    `json:"role_id"`
	Level   int       `json:"level"`
	Expire  time.Time `json:"expire"`
}
