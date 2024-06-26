package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type SysGroup struct {
	common.BaseModel
	ShortName string    `json:"shortName" gorm:"index;not null;comment:组织简称"`
	OrgCode   string    `json:"orgCode" gorm:"uniqueIndex;not null;comment:组织编码"`
	ParentID  uint      `json:"parentID" gorm:"index;default:0;comment:父ID"`
	Sort      uint8     `json:"sort" gorm:"index;default:0;comment:排序"`
	Enable    bool      `json:"enable" gorm:"index;default:true;comment:是否可用"`
	SysUsers  []SysUser `json:"sysUsers" gorm:"many2many:sys_user_group;"`
	SysRoles  []SysRole `json:"sysRoles"`
}

func (s *SysGroup) TableName() string {
	return "sys_group"
}

var sysGroupWorkerId int64 = global.SysGroupWorkerId

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *SysGroup) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(sysGroupWorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
