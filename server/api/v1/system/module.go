package system

import "github.com/jasvtfvan/oms-admin/server/service"

type ApiGroup struct {
	DbApi
}

var (
	initDBService   = service.ServiceGroupApp.Initialize.InitDBService
	updateDBService = service.ServiceGroupApp.Initialize.UpdateDBService
)
