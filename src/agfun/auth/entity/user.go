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
	RoleID      uint   `json:"role_id"`
}
type VipUser struct {
	gorm.Model
	UserID uint      `json:"user_id"`
	Level  int       `json:"level"`
	Expire time.Time `json:"expire"`
}
type UserRole struct {
	gorm.Model
	Name  string `json:"name"`
	Level int    `json:"level"`
}
type WXFriend struct {
	gorm.Model
	UserName    string `json:"user_name"`
	NickName    string `json:"nick_name"`
	HeadImgUrl  string `json:"head_img_url"`
	ContactFlag int    `json:"contact_flag"`
	MemberCount int    `json:"member_count"`
	Sex         int    `json:"sex"`
	VerifyFlag  int    `json:"verify_flag"` //验证标志
	Province    string `json:"province"`    //省
	City        string `json:"city"`        //市
}
