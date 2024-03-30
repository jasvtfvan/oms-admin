package system

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
