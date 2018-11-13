package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName    string `json:"user_name"`
	Pwd         string `json:"pwd"`
	NickName    string
	AccessToken string
}
