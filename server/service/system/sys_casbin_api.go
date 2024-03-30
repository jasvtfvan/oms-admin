package system

import "github.com/jasvtfvan/oms-admin/server/model/system"

type CasbinApiService interface {
	GetCasbinApiList() []system.SysCasbinApi
}

type CasbinApiServiceImpl struct{}

func (*CasbinApiServiceImpl) GetCasbinApiList() []system.SysCasbinApi {
	var res = []system.SysCasbinApi{
		{Path: "/user/delete", Method: "DELETE"},
		{Path: "/user/disable", Method: "PUT"},
		{Path: "/user/enable", Method: "PUT"},
		{Path: "/user/reset-pwd", Method: "PUT"},
	}
	return res
}
