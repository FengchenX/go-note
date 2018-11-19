package service

import (
	"agfun/agfun-service/dbcentral/etcddb"
	"agfun/agfun-service/dbcentral/mysqldb"
	"github.com/jinzhu/gorm"
)

type Svc struct {
	Dynamic *etcddb.Client
	SysDB   *gorm.DB
	AuthDB  *gorm.DB
}

func NewSvc(dynamic *etcddb.Client, sys, auth *gorm.DB) *Svc {
	return &Svc{
		Dynamic: dynamic,
		SysDB:   sys,
		AuthDB:  auth,
	}
}

var std *Svc

func Init() {
	if std != nil {
		return
	}
	etcddb.Init()
	mysqldb.Init()
	std = NewSvc(etcddb.GetCli(), mysqldb.GetSysDB(), mysqldb.GetAuthDB())
}

func GetDefaultSvc() *Svc {
	return std
}
