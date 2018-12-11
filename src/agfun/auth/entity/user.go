package entity

import (
	"time"
)

// User represents the user for this application
//
// swagger:model
type User struct {
	ID          string `json:"id"`
	UserName    string `json:"user_name"`
	Pwd         string `json:"pwd"`
	NickName    string `json:"nick_name"`
	WeChat      string `json:"we_chat"`
	AccessToken string `json:"access_token"`
	RoleID      uint   `json:"role_id"`
}
// swagger:model
type VipUser struct {
	ID     string    `json:"id"`
	UserID uint      `json:"user_id"`
	Level  int       `json:"level"`
	Expire time.Time `json:"expire"`
}
// swagger:model
type UserRole struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
// swagger:model
type WXFriend struct {
	ID          string `json:"id"`
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
