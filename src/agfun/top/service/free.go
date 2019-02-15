package service

import (
	"agfun/dto"
	"agfun/top/dbcentral/mysqldb"
	"agfun/top/entity"
)

func (s *TVSvc) AddFreeTVs(frees []*entity.FreeTV) error {
	for _, free := range frees {
		_, i, e := mysqldb.GetFreeTVs(*free, nil)
		if e != nil {
			return e
		}
		if i > 0 {
			continue
		}
		e = mysqldb.AddFreeTVs(free)
		if e != nil {
			return e
		}
	}
	return nil
}
func (s *TVSvc) GetFreeTVs(req dto.GetVideos) (*entity.GetTVResp, error) {
	movies, i, e := mysqldb.GetFreeTVs(entity.FreeTV{}, req.Filter)
	if e != nil {
		return nil, e
	}
	var resp entity.GetTVResp
	resp.Total = i
	resp.Frees = movies
	return &resp, nil
}

func (s *TVSvc) UpdateFreeTV(free entity.FreeTV) error {
	if len(free.ID) > 0 {
		e := mysqldb.UpdateFreeTV(free, nil)
		if e != nil {
			return e
		}
		return nil
	}
	//todo
	return nil
}

func (s *TVSvc) DelFreeTV(free entity.FreeTV, token string) error {
	e := mysqldb.DelFreeTV(free)
	return e
}
