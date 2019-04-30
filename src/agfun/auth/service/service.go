package service

import (
	"agfun/auth/dbcentral/etcd"
	"agfun/service"
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
