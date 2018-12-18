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

func initStdSvc() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewAuthSvc()

	stdSvc.Svc = service.GetDefaultSvc()

	stdSvc.Dynamic = etcd.GetDefaultCli()
}
func GetDefaultSvc() *AuthSvc {
	if stdSvc == nil {
		initStdSvc()
	}
	return stdSvc
}
