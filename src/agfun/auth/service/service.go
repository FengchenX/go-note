package service

import (
	"agfun/agfun-service/service"
	"agfun/auth/dbcentral/etcd"
)

type AuthSvc struct {
	*service.Svc
	Dynamic *etcd.Client
}

func NewAuthSvc() *AuthSvc {
	return &AuthSvc{}
}

var stdSvc *AuthSvc

func Init() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewAuthSvc()

	service.Init()
	stdSvc.Svc = service.GetDefaultSvc()

	etcd.Init()
	stdSvc.Dynamic = etcd.GetDefaultCli()
}
func GetDefaultSvc() *AuthSvc {
	return stdSvc
}
