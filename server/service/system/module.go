package system

type ServiceGroup struct {
	DbService
	DemoService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		DbService:   new(DbServiceImpl),
		DemoService: new(DemoServiceImpl),
	}
	return group
}