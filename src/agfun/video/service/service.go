package service

import (
	"agfun/video/dbcentral/etcd"
	"agfun/video/dbcentral/pg"
	etcd2 "db/etcd"
	pg2 "db/pg"
	"service"
)

type VideoSvc struct {
	service.Svc
	SysDB   *pg.SysDB
	AuthDB  *pg.AuthDB
	Dynamic *etcd.Client
}

func NewMovieSvc() *VideoSvc {
	svc := &VideoSvc{
		Svc:     *service.DefaultSvc(),
		SysDB:   &pg.SysDB{SysDB: *pg2.DefaultSysDB()},
		AuthDB:  &pg.AuthDB{AuthDB: *pg2.DefaultAuthDB()},
		Dynamic: &etcd.Client{Client: *etcd2.DefaultCli()},
	}
	return svc
}
