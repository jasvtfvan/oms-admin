package service

import (
	"github.com/jasvtfvan/oms-admin/server/service/goods"
	"github.com/jasvtfvan/oms-admin/server/service/system"
)

type ServiceGroup struct {
	System system.ServiceGroup
	Goods  goods.ServiceGroup
}

var ServiceGroupApp = &ServiceGroup{
	System: *system.NewServiceGroup(),
	Goods:  *goods.NewServiceGroup(),
}
