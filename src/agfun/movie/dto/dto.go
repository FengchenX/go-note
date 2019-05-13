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

type Movies struct {
	Total int `json:"total"`
	Movies []Movie `json:"movies"`
}

type MovieVideos struct {
	Total int `json:"total"`
	MovieVideos []MovieVideo `json:"movie_videos"`
}
