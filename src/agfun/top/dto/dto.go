package dto

import "agfun/entity"

type Banner struct {
	ID string `json:"id"`
	Video entity.Video `json:"video"`
}

type GetBannersResp struct {
	Count int `json:"count"`
	Data []Banner `json:"data"`
}