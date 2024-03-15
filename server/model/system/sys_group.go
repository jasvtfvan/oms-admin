package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
)

type SysGroup struct {
	common.BaseModel
	ShortName string    `json:"shortName" gorm:"index;not null;comment:组织简称"`
	OrgCode   string    `json:"orgCode" gorm:"index;not null;comment:组织编码"`
	ParentID  uint      `json:"parentID" gorm:"default:0;comment:父ID"`
	Sort      uint8     `json:"sort" gorm:"default:0;comment:排序"`
	Enable    bool      `json:"enable" gorm:"default:true;comment:是否可用"`
	SysUsers  []SysUser `gorm:"many2many:sys_user_group;"`
	SysRoles  []SysRole
}

func (SysGroup) TableName() string {
	return "sys_group"
}
