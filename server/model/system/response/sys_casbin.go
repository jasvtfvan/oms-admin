package response

type CasbinInfo struct {
	Path    string `json:"path"`    // 路径obj
	Method  string `json:"method"`  // 方法act
	Checked bool   `json:"checked"` // 是否已选中（即拥有该权限）
}

// 需要通过casbin验证的所有资源
func DefaultCasbinSource() []CasbinInfo {
	btns := DefaultCasbinBtn()
	apis := DefaultCasbinApi()
	target := []CasbinInfo{}
	target = append(target, btns...)
	target = append(target, apis...)
	return target
}

// 需要通过casbin验证的空资源
func DefaultCasbinBtn() []CasbinInfo {
	return []CasbinInfo{}
}

// 需要通过casbin验证的api
func DefaultCasbinApi() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/user/delete", Method: "DELETE", Checked: false},
		{Path: "/user/disable", Method: "PUT", Checked: false},
		{Path: "/user/enable", Method: "PUT", Checked: false},
		{Path: "/user/reset-pwd", Method: "PUT", Checked: false},
	}
}
