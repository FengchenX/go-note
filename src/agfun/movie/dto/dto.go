package dto

import (
	"agfun/movie/entity"
)

type Movie struct {
	entity.Movie
	MainPlayers []string `json:"main_players"`
	Types       []string `json:"types"`
}

type MovieVideo struct {
	entity.MovieVideo
}

//type MovieVideos []MovieVideo

type Movies struct {
	Total int `json:"total"`
	Movies []Movie `json:"movies"`
}
