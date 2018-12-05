package service

import "agfun/agfun-service/service"

type TVSvc struct {
	*service.Svc
}

func NewTVSvc() *TVSvc {
	return &TVSvc{}
}

var stdSvc *TVSvc

func initStdSvc() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewTVSvc()
	stdSvc.Svc = service.GetDefaultSvc()
}
func GetDefaultSvc() *TVSvc {
	if stdSvc == nil {
		initStdSvc()
	}
	return stdSvc
}
