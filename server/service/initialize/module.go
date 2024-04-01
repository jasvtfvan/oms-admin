package initialize

import freecache "github.com/jasvtfvan/oms-admin/server/utils/freecache"

var cacheStore = freecache.GetStoreDefault()

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
