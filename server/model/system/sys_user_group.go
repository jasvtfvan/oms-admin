package system

type SysUserGroup struct {
	SysUserID  uint `json:"sysUserID" gorm:"primaryKey"`
	SysGroupID uint `json:"sysGroupID" gorm:"primaryKey"`
}

func (s *SysUserGroup) TableName() string {
	return "sys_user_group"
}
