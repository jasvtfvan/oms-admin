package system

type ServiceGroup struct {
	DBService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		DBService: new(DBServiceImpl),
	}
	return group
}
