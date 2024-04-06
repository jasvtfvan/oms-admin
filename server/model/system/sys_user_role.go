package system

type SysUserRole struct {
	SysUserID uint `json:"sysUserId" gorm:"primaryKey"`
	SysRoleID uint `json:"sysRoleId" gorm:"primaryKey"`
}

func (s *SysUserRole) TableName() string {
	return "sys_user_role"
}
