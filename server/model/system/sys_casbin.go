package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type SysCasbin struct {
	ID    uint   `gorm:"primaryKey"`
	Ptype string `gorm:"size:100"`
	V0    string `gorm:"size:100"`
	V1    string `gorm:"size:100"`
	V2    string `gorm:"size:100"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}

func (s *SysCasbin) TableName() string {
	return "sys_casbin"
}

var sysCasbinWorkerId int64 = global.SysCasbinWorkerId

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *SysCasbin) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(sysCasbinWorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
