package dto

import (
	"agfun/movie/entity"
	. "dto"
)

type Movie struct {
	entity.Movie
	Video Video `json:"video"`
}

type FreeMovie struct {
	entity.FreeMovie
	Movie Movie `json:"movie"`
}

type FreeMovies struct {
	Total int `json:"total"`
	Rows []FreeMovie `json:"rows"`
}
