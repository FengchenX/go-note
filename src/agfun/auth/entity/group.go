package entity

import "time"

// group
// swagger:model
type Group struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Max        int       `json:"max"`
	ParentID   string    `json:"parent_id"`
	CreateTime time.Time `json:"create_time"`
}

// group user
// swagger:model
type GroupUser struct {
	GroupID    string    `json:"group_id"`
	UserID     string    `json:"user_id"`
	CreateTime time.Time `json:"create_time"`
}
