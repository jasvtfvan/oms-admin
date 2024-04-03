package system

import "github.com/jasvtfvan/oms-admin/server/dao"

type ServiceGroup struct {
	CasbinApiService
	CasbinService
	GroupService
	JWTService
	OperationRecordService
	UserService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		CasbinApiService:       new(CasbinApiServiceImpl),
		CasbinService:          new(CasbinServiceImpl),
		GroupService:           new(GroupServiceImpl),
		JWTService:             new(JWTServiceImpl),
		OperationRecordService: new(OperationRecordServiceImpl),
		UserService:            new(UserServiceImpl),
	}
	return group
}

var (
	CasbinDao = dao.DaoGroupApp.System.CasbinDao
)
