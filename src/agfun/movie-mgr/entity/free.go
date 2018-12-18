package entity

import "agfun/agfun-service/entity"

// free movie
// swagger:model
type FreeMovie struct {
	Movie
	entity.FreeVideo
}

type GetMoviesResp struct {
	Total int
	Frees []*FreeMovie
}
