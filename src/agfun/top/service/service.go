package service

import "agfun/service"

type TopSvc struct {
	*service.Svc
}

func NewTopSvc() *TopSvc {
	return &TopSvc{}
}

var stdSvc *TopSvc

func initStdSvc() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewTopSvc()
	stdSvc.Svc = service.GetDefaultSvc()
}
func GetDefaultSvc() *TopSvc {
	if stdSvc == nil {
		initStdSvc()
	}
	return stdSvc
}
