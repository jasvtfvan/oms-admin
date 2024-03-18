package goods

type ServiceGroup struct {
	DemoService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		DemoService: new(DemoServiceImpl),
	}
	return group
}
