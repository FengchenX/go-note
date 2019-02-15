package entity

import "agfun/entity"

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
