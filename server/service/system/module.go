package system

import jwtRedis "github.com/jasvtfvan/oms-admin/server/utils/redis/jwt"

var jwtStore = jwtRedis.GetRedisStore()

type ServiceGroup struct {
	UserService
	JWTService
	CasbinApiService
	OperationRecordService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		UserService:            new(UserServiceImpl),
		JWTService:             new(JWTServiceImpl),
		CasbinApiService:       new(CasbinApiServiceImpl),
		OperationRecordService: new(OperationRecordServiceImpl),
	}
	return group
}
