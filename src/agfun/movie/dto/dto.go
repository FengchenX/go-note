package dto

import (
	"agfun/movie/entity"
)

type Movie struct {
	entity.Movie
	MainPlayers []string `json:"main_players"`
	Types       []string `json:"types"`
}
