package entity

import "agfun/entity"

type FreeTV struct {
	TV
	entity.FreeVideo
}

type GetTVResp struct {
	Total int
	Frees []*FreeTV
}
