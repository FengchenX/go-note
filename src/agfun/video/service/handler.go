package service

import (
	"agfun/video/dto"
	"agfun/video/entity"
	"github.com/kataras/iris"
	"util"
)

func (s *VideoSvc) AddVideo(c iris.Context) {
	req := dto.Video{}
	e := c.ReadJSON(&req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	m := entity.Video{}
	util.Copy(&m, &req.Video)
	m.ID = util.NewUUID()
	m.CreateAt = util.TimeNowStd()
	e = s.SysDB.AddVideo(&m)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func (s *VideoSvc) GetVideos(c iris.Context) {
	req := dto.Video{}
	filter := util.ParsePageFilter(c)
	videos, i, e := s.SysDB.GetVideos(&req.Video, filter)
	if e != nil {
		util.Fail(c, e)
		return
	}
	res := dto.Videos{Total: i}
	for _, video := range videos {
		dtoVideo := dto.Video{
			Video: video,
		}
		res.Videos = append(res.Videos, dtoVideo)
	}
	util.Success(c, &res)
}
func (s *VideoSvc) GetVideo(c iris.Context) {
	id:=c.Params().Get("id")
	v:=entity.Video{ID:id}
	videos, _, e := s.SysDB.GetVideos(&v, nil)
	if e != nil {
		util.Fail(c, e)
		return
	}
	res:=dto.Video{Video:videos[0]}
	util.Success(c, &res)
}
func (s *VideoSvc) UpVideo(c iris.Context) {

}
func (s *VideoSvc) DelVideo(c iris.Context) {

}
