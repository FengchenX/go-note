package dto

import (
	"agfun/movie-mgr/entity"
	"net/http"
)

// get movies
// swagger:parameters getFreeMovies
type GetFreeMoviesParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
	// auth session
	// in: header
	Session string `json:"session"`
	// test
	// in: query
	Test string
}

// add movies
// swagger:parameters addFreeMovies
type AddFreeMoviesParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
	// auth session
	// in: header
	Session string `json:"session"`
	
	// in:body
	Body entity.FreeMovie `json:"body"`
}
