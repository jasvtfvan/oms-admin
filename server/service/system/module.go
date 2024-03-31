package system

import jwtRedis "github.com/jasvtfvan/oms-admin/server/utils/redis/jwt"

var jwtStore = jwtRedis.GetRedisStore()

type ServiceGroup struct {
	UserService
	JWTService
	CasbinApiService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		UserService:      new(UserServiceImpl),
		JWTService:       new(JWTServiceImpl),
		CasbinApiService: new(CasbinApiServiceImpl),
	}
	return group
}
