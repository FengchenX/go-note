package service

import (
	"agfun/movie/dbcentral/etcd"
	"agfun/movie/dbcentral/pg"
	etcd2 "db/etcd"
	pg2 "db/pg"
	"service"
)

type MovieSvc struct {
	service.Svc
	SysDB   *pg.SysDB
	AuthDB  *pg.AuthDB
	Dynamic *etcd.Client
}

func NewMovieSvc() *MovieSvc {
	svc := &MovieSvc{
		Svc:     *service.DefaultSvc(),
		SysDB:   &pg.SysDB{SysDB: *pg2.DefaultSysDB()},
		AuthDB:  &pg.AuthDB{AuthDB: *pg2.DefaultAuthDB()},
		Dynamic: &etcd.Client{Client: *etcd2.DefaultCli()},
	}
	return svc
}
