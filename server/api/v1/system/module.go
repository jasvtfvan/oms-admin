package system

import "github.com/jasvtfvan/oms-admin/server/service"

type ApiGroup struct {
	DbApi
	UserApi
	CasbinApi
	CacheApi
}

var (
	initDBService    = service.ServiceGroupApp.Initialize.InitDBService
	updateDBService  = service.ServiceGroupApp.Initialize.UpdateDBService
	userService      = service.ServiceGroupApp.System.UserService
	jwtService       = service.ServiceGroupApp.System.JWTService
	casbinApiService = service.ServiceGroupApp.System.CasbinApiService
	groupService     = service.ServiceGroupApp.System.GroupService
)
