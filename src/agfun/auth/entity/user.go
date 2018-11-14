package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName    string `json:"user_name"`
	Pwd         string `json:"pwd"`
	NickName    string `json:"nick_name"`
	AccessToken string `json:"access_token"`
}
