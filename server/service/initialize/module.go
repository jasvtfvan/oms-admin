package initialize

type ServiceGroup struct {
	InitDBService
	UpdateDBService
}

func NewServiceGroup() *ServiceGroup {
	group := &ServiceGroup{
		InitDBService:   new(InitDBServiceImpl),
		UpdateDBService: new(UpdateDBServiceImpl),
	}
	return group
}
