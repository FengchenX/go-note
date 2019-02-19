package service

import (
	"agfun/entity"
	"agfun/top/dto"
)

func (s *TopSvc) GetBanners() (dto.GetBannersResp, error) {
	//panic("todo")
	return dto.GetBannersResp{
		Count: 2,
		Data:  []dto.Banner{{
			ID:    "",
			Video: entity.Video{
				ID:   "111",
				Name: "name1",
				Pic:  "http://i0.hdslb.com/bfs/archive/5c1ad0c27885c11d9de7729abd3dd20d923ce745.jpg",
				Desc: "",
				URL:  "",
			},
		}, {
			ID:    "",
			Video: entity.Video{
				ID:   "2222",
				Name: "name2",
				Pic:  "http://i0.hdslb.com/bfs/archive/e20e2e856f7bad9123b21168dd1b876676887b4e.png",
				Desc: "",
				URL:  "",
			},
		}},
	}, nil
}

