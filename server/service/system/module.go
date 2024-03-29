package system

type ServiceGroup struct {
	UserService
	JWTService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		UserService: new(UserServiceImpl),
		JWTService:  new(JWTServiceImpl),
	}
	return group
}
