package response

type ProfileRole struct {
	RoleName string `json:"roleName"` // 角色名
	RoleCode string `json:"roleCode"` // 角色编码（唯一）
	IsAdmin  bool   `json:"isAdmin"`  // 是否管理员
	Sort     uint8  `json:"sort"`     // 排序字段
}

type ProfileGroup struct {
	ShortName string        `json:"shortName"` // 组织简称
	OrgCode   string        `json:"orgCode"`   // 组织编码（唯一）
	Sort      uint8         `json:"sort"`      // 排序
	SysRoles  []ProfileRole `json:"sysRoles"`  // 组织下的用户绑定的角色
}

type UserProfile struct {
	Username     string         `json:"username"`     // 用户名
	NickName     string         `json:"nickName"`     // 昵称
	Avatar       string         `json:"avatar"`       // 头像
	Phone        string         `json:"phone"`        // 手机号
	Email        string         `json:"email"`        // 邮箱
	LogOperation bool           `json:"logOperation"` // 是否记录操作记录
	Enable       bool           `json:"enable"`       // 是否可用
	IsRootAdmin  bool           `json:"isRootAdmin"`  // 是否系统管理员
	SysGroups    []ProfileGroup `json:"sysGroups"`    // 关联的组织
}
