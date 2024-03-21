package initialize

type ServiceGroup struct {
	InitDBService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		InitDBService: new(InitDBServiceImpl),
	}
	return group
}
