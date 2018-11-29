package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName    string `json:"user_name"`
	Pwd         string `json:"pwd"`
	NickName    string `json:"nick_name"`
	WeChat      string `json:"we_chat"`
	AccessToken string `json:"access_token"`
}
type VipUser struct {
	gorm.Model
	UserID uint      `json:"user_id"`
	Level  int       `json:"level"`
	Expire time.Time `json:"expire"`
}
