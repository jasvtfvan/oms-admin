package system

import "github.com/jasvtfvan/oms-admin/server/model/system"

type CasbinApi struct{}

func (*CasbinApi) GetCasbinApiList() []system.SysCasbinApi {
	return casbinApiService.GetCasbinApiList()
}
