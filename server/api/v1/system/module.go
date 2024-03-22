package system

import "github.com/jasvtfvan/oms-admin/server/service"

type ApiGroup struct {
	DbApi
}

var (
	systemDBService = service.ServiceGroupApp.System.DBService
	initDBService   = service.ServiceGroupApp.Initialize.InitDBService
	updateDBService = service.ServiceGroupApp.Initialize.UpdateDBService
)
