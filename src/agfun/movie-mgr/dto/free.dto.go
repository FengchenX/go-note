package dto

import "agfun/agfun-service/util"

type GetFreeVideos struct {
	Token  string
	Filter *util.PageFilter
}
