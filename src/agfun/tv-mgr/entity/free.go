package entity

import "agfun/agfun-service/entity"

type FreeTV struct {
	TV
	entity.FreeVideo
}

type GetTVResp struct {
	Total int
	Frees []*FreeTV
}
