package entity

import "time"

type User struct {
	Id uint
	UserName string
	Pwd string
	NickName string
	CreateAt time.Time
	AccessToken string
}
