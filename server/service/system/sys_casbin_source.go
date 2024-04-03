package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/system/response"
)

type CasbinApiService interface {
	GetCasbinApiList() []response.CasbinInfo
}

type CasbinApiServiceImpl struct{}

func (*CasbinApiServiceImpl) GetCasbinApiList() []response.CasbinInfo {
	return response.DefaultCasbin()
}
