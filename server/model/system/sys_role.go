package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
)

type SysRole struct {
	common.BaseModel
	RoleName   string `json:"roleName" gorm:"index;not null;comment:角色名称"`
	RoleCode   string `json:"roleCode" gorm:"index;not null;comment:角色编码"`
	Sort       uint8  `json:"sort" gorm:"default:0;comment:排序"`
	Comment    string `json:"comment" gorm:"default:'';comment:备注"`
	Enable     bool   `json:"enable" gorm:"default:true;comment:是否可用"`
	SysGroupID uint
	SysUsers   []SysUser `gorm:"many2many:sys_user_role;"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
