package service

import (
	"github.com/jasvtfvan/oms-admin/server/service/demo"
	"github.com/jasvtfvan/oms-admin/server/service/system"
)

type ServiceGroup struct {
	System system.ServiceGroup
	Demo   demo.ServiceGroup
}

var ServiceGroupApp = &ServiceGroup{
	System: *system.NewServiceGroup(),
	Demo:   *demo.NewServiceGroup(),
}
