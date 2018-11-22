package dto

import "agfun/agfun-service/util"

type GetVideos struct {
	Token  string
	Filter *util.PageFilter
}
