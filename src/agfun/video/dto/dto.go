package dto

import "agfun/video/entity"

type Video struct {
	entity.Video
}

type Videos struct {
	Total  int     `json:"total"`
	Videos []Video `json:"videos"`
}
