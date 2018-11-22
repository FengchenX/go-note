package dto

import "agfun/agfun-service/util"

type GetFreeMovies struct {
	Token  string
	Filter *util.PageFilter
}
