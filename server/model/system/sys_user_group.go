package system

type SysUserGroup struct {
	SysUserID  uint `json:"sysUserId" gorm:"primaryKey"`
	SysGroupID uint `json:"sysGroupId" gorm:"primaryKey"`
}

func (s *SysUserGroup) TableName() string {
	return "sys_user_group"
}
