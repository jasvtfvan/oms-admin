package system

type ServiceGroup struct {
	UserService
	JWTService
	CasbinApiService
	OperationRecordService
	GroupService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		UserService:            new(UserServiceImpl),
		JWTService:             new(JWTServiceImpl),
		CasbinApiService:       new(CasbinApiServiceImpl),
		OperationRecordService: new(OperationRecordServiceImpl),
		GroupService:           new(GroupServiceImpl),
	}
	return group
}
