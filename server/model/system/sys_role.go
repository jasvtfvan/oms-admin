package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type SysRole struct {
	common.BaseModel
	RoleName   string    `json:"roleName" gorm:"index;not null;comment:角色名称"`
	RoleCode   string    `json:"roleCode" gorm:"index;not null;comment:角色编码"`
	Sort       uint8     `json:"sort" gorm:"index;default:0;comment:排序"`
	Comment    string    `json:"comment" gorm:"default:'';comment:备注"`
	Enable     bool      `json:"enable" gorm:"index;default:true;comment:是否可用"`
	SysGroupID uint      `json:"sysGroupID" gorm:"index;not null;comment:组织ID"`
	SysUsers   []SysUser `gorm:"many2many:sys_user_role;"`
}

func (*SysRole) TableName() string {
	return "sys_role"
}

var sysRoleWorkerId int64 = sysGroupWorkerId + 1

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *SysRole) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(sysRoleWorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
