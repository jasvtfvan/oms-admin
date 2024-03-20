package system

type SysUserRole struct {
	SysUserID uint `json:"sysUserID" gorm:"primaryKey"`
	SysRoleID uint `json:"sysRoleID" gorm:"primaryKey"`
}

func (s *SysUserRole) TableName() string {
	return "sys_user_role"
}
