package service

import (
	"agfun/movie-mgr/dbcentral/mysqldb"
	"agfun/movie-mgr/dto"
	"agfun/movie-mgr/entity"
)

func (s *MgrSvc) AddFreeVideos(frees []*entity.FreeVideo) error {
	for _, free := range frees {
		_, i, e := mysqldb.GetFreeVideos(*free, nil)
		if e != nil {
			return e
		}
		if i > 0 {
			continue
		}
		e = mysqldb.AddFreeVideo(free)
		if e != nil {
			return e
		}
	}
	return nil
}
func (s *MgrSvc) GetFreeVideos(req dto.GetFreeVideos) ([]*entity.FreeVideo, error) {
	var id int
	e := s.Dynamic.Get(req.Token, &id)
	if e != nil {
		return nil, e
	}
	var videos []*entity.FreeVideo
	return videos, nil
}
