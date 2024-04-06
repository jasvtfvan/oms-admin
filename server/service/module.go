package service

import (
	"github.com/jasvtfvan/oms-admin/server/service/initialize"
	"github.com/jasvtfvan/oms-admin/server/service/system"
)

type ServiceGroup struct {
	Initialize initialize.ServiceGroup
	System     system.ServiceGroup
}

var ServiceGroupApp = &ServiceGroup{
	Initialize: *initialize.NewServiceGroup(),
	System:     *system.NewServiceGroup(),
}
