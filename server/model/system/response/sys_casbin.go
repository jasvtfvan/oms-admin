package response

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/user/delete", Method: "DELETE"},
		{Path: "/user/disable", Method: "PUT"},
		{Path: "/user/enable", Method: "PUT"},
		{Path: "/user/reset-pwd", Method: "PUT"},
	}
}
