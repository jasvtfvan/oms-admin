package response

type LoginRole struct {
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
	Sort     uint8  `json:"sort"`
}

type LoginGroups struct {
	ShortName string      `json:"shortName"`
	OrgCode   string      `json:"orgCode"`
	Sort      uint8       `json:"sort"`
	SysRoles  []LoginRole `json:"sysRoles"`
}

type LoginUser struct {
	Username     string        `json:"username"`
	NickName     string        `json:"nickName"`
	Avatar       string        `json:"avatar"`
	Phone        string        `json:"phone"`
	Email        string        `json:"email"`
	IsAdmin      bool          `json:"isAdmin"`
	LogOperation bool          `json:"logOperation"`
	SysGroups    []LoginGroups `json:"sysGroups"`
}

type Login struct {
	User  LoginUser `json:"user"`
	Token string    `json:"token"`
}
