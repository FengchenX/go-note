package service

import (
	mysqldb2 "agfun/dbcentral/mysqldb"
	"agfun/entity"
	"agfun/top/conf"
	"agfun/top/dbcentral/mysqldb"
	"agfun/top/dto"
	entity2 "agfun/top/entity"
	"agfun/util"
	"fmt"
	"time"
)

func (s *TopSvc) GetBanners() (dto.GetBannersResp, error) {
	temp, e := time.Parse("2006-01-02 03:04:05", conf.TopConfInst().TopCreateTime)
	if e != nil {
		return dto.GetBannersResp{}, e
	}
	banner := entity2.Banner{
		CreateTime: temp,
	}
	banners, e := mysqldb.GetBanners(banner)
	var resp dto.GetBannersResp
	resp.Count = len(banners)
	for _, banner := range banners {
		video := entity.Video{
			ID: banner.VideoID,
		}
		videos, i, e := mysqldb2.GetVideos(video, nil)
		if e != nil {
			return dto.GetBannersResp{}, e
		}
		if i != 1 {
			return dto.GetBannersResp{}, fmt.Errorf("is != 1")
		}
		dtoVideo := dto.Banner{
			ID:    banner.ID,
			Video: entity.Video{},
		}
		e = util.Copy(&dtoVideo.Video, &videos[0])
		if e != nil {
			return dto.GetBannersResp{}, e
		}
		resp.Data = append(resp.Data, dtoVideo)
	}
	return resp, e
}
