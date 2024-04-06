package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/system/response"
)

type CasbinApiService interface {
	GetCasbinApiList(roleCode, groupCode string) []response.CasbinInfo
}

type CasbinApiServiceImpl struct{}

// 根据role和group获取所有api列表，包含是否已经选中各api的权限
func (*CasbinApiServiceImpl) GetCasbinApiList(roleCode, groupCode string) []response.CasbinInfo {
	rules := CasbinServiceApp.GetPolicyInEnforcer(roleCode, groupCode)
	casbinReqInfos := CasbinServiceApp.GetCasbinInfoByPolicy(rules)
	casbinResInfos := response.DefaultCasbinApi()
	result := []response.CasbinInfo{}
	for _, res := range casbinResInfos { // 循环所有api
		for _, req := range casbinReqInfos { // 循环具有权限的api
			// 找到权限
			if req.Path == res.Path && req.Method == res.Method {
				result = append(result, response.CasbinInfo{
					Path:    req.Path,
					Method:  req.Method,
					Checked: true,
				})
			}
		}
		// 循环后未找到权限
		result = append(result, response.CasbinInfo{
			Path:    res.Path,
			Method:  res.Method,
			Checked: false,
		})
	}
	return result
}
