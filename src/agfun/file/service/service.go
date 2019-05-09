package service

import (
	"agfun/file/dbcentral/etcd"
	"agfun/file/dbcentral/pg"
	"service"
)

type FileSvc struct {
	service.Svc
	SysDB   *pg.SysDB
	AuthDB  *pg.AuthDB
	Dynamic *etcd.Client
}

func NewFileSvc() *FileSvc {
	svc := &FileSvc{
		Svc:     *service.DefaultSvc(),
		//SysDB:   &pg.SysDB{SysDB: *pg2.DefaultSysDB()},
		//AuthDB:  &pg.AuthDB{AuthDB: *pg2.DefaultAuthDB()},
		//Dynamic: &etcd.Client{Client: *etcd2.DefaultCli()},
	}
	return svc
}