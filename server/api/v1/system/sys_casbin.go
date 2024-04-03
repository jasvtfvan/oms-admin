package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/system/response"
)

type CasbinApi struct{}

func (*CasbinApi) GetCasbinApiList() []response.CasbinInfo {
	return casbinApiService.GetCasbinApiList()
}
