package dto

import (
	entity2 "agfun/entity"
	"agfun/movie/entity"
)

//type GetMovies struct {
//	Token  string
//	Filter *util.PageFilter
//}

type FreeMovie struct {
	entity.Movie
	entity2.FreeVideo
	entity2.Video
}

type GetMoviesResp struct {
	Total int          `json:"total"`
	Frees []*FreeMovie `json:"frees"`
}
