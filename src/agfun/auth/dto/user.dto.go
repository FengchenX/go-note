package dto

import "agfun/auth/entity"

type VipUser struct {
	entity.VipUser
	Expire string `json:"expire"`
}
