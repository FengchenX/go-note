package dto

import (
	"agfun/movie/entity"
	. "dto"
)

type Movie struct {
	entity.Movie
	Video Video `json:"video"`
}
