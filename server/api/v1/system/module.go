package system

import "github.com/jasvtfvan/oms-admin/server/service"

type ApiGroup struct {
	CacheApi
	CasbinApi
	DbApi
	UserApi
}

var (
	casbinApiService = service.ServiceGroupApp.System.CasbinApiService
	groupService     = service.ServiceGroupApp.System.GroupService
	initDBService    = service.ServiceGroupApp.Initialize.InitDBService
	jwtService       = service.ServiceGroupApp.System.JWTService
	roleService      = service.ServiceGroupApp.System.RoleService
	updateDBService  = service.ServiceGroupApp.Initialize.UpdateDBService
	userService      = service.ServiceGroupApp.System.UserService
)
