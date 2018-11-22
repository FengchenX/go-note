package entity

import "agfun/agfun-service/entity"

type FreeMovie struct {
	Movie
	entity.FreeVideo
}

type GetMoviesResp struct {
	Total int
	Frees []*FreeMovie
}
