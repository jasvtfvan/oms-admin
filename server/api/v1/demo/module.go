package demo

import "github.com/jasvtfvan/oms-admin/server/service"

type ApiGroup struct {
	DemoApi
}

var (
	demoService = service.ServiceGroupApp.Demo.DemoService
)
