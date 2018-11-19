package service

import "agfun/agfun-service/service"

type MgrSvc struct {
	*service.Svc
}

func NewAuthSvc() *MgrSvc {
	return &MgrSvc{}
}

var stdSvc *MgrSvc

func initStdSvc() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewAuthSvc()
	stdSvc.Svc = service.GetDefaultSvc()
}
func GetDefaultSvc() *MgrSvc {
	if stdSvc == nil {
		initStdSvc()
	}
	return stdSvc
}