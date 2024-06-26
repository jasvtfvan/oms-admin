package system

import "github.com/jasvtfvan/oms-admin/server/dao"

type ServiceGroup struct {
	CasbinApiService
	CasbinService
	GroupService
	JWTService
	OperationRecordService
	RoleService
	UserService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		CasbinApiService:       new(CasbinApiServiceImpl),
		CasbinService:          new(CasbinServiceImpl),
		GroupService:           new(GroupServiceImpl),
		JWTService:             new(JWTServiceImpl),
		OperationRecordService: new(OperationRecordServiceImpl),
		RoleService:            new(RoleServiceImpl),
		UserService:            new(UserServiceImpl),
	}
	return group
}

var (
	casbinDao          = dao.DaoGroupApp.System.CasbinDao
	groupDao           = dao.DaoGroupApp.System.GroupDao
	operationRecordDao = dao.DaoGroupApp.System.OperationRecordDao
	roleDao            = dao.DaoGroupApp.System.RoleDao
	userDao            = dao.DaoGroupApp.System.UserDao
	userGroupDao       = dao.DaoGroupApp.System.UserGroupDao
	userRoleDao        = dao.DaoGroupApp.System.UserRoleDao
)
