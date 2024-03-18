package system

import "github.com/jasvtfvan/oms-admin/server/service"

type ApiGroup struct {
	DbApi
}

var (
	systemDbService = service.ServiceGroupApp.System.DbService
)
