package entity

import (
	"time"
)

// User represents the user for this application
//
// swagger:model
type User struct {
	ID          string    `json:"id"`
	UserName    string    `json:"user_name"`
	Pwd         string    `json:"pwd"`
	NickName    string    `json:"nick_name"`
	WeChat      string    `json:"we_chat"`
	AccessToken string    `json:"access_token"`
	RoleID      uint      `json:"role_id"`
	CreateTime  time.Time `json:"create_time"`
}

// friend
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
