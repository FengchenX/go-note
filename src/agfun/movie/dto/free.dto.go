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
	Movie entity.Movie `json:"movie"`
	FreeVideo entity2.FreeVideo `json:"free_video"`
	Video entity2.Video `json:"video"`
}

type GetMoviesResp struct {
	Total int          `json:"total"`
	Frees []*FreeMovie `json:"frees"`
}
