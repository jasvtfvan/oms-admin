package service

import (
	"github.com/jasvtfvan/oms-admin/server/service/demo"
	"github.com/jasvtfvan/oms-admin/server/service/initialize"
	"github.com/jasvtfvan/oms-admin/server/service/system"
)

type ServiceGroup struct {
	Initialize initialize.ServiceGroup
	System     system.ServiceGroup
	Demo       demo.ServiceGroup
}

var ServiceGroupApp = &ServiceGroup{
	Initialize: *initialize.NewServiceGroup(),
	System:     *system.NewServiceGroup(),
	Demo:       *demo.NewServiceGroup(),
}
