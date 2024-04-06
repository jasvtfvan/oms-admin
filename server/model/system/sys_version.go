package system

import (
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type SysVersion struct {
	ID          uint      `json:"ID" gorm:"primaryKey"`
	VersionCode string    `json:"versionCode" gorm:"uniqueIndex;not null;comment:版本编号"`
	VersionName string    `json:"versionName" gorm:"index;not null;comment:版本名称"`
	Version     string    `json:"version" gorm:"default:0.0.1;comment:版本号"`
	CreatedAt   time.Time `json:"createdAt"` // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"` // 更新时间
}

func (s *SysVersion) TableName() string {
	return "sys_version"
}

var sysVersionWorkerId int64 = global.SysVersionWorkerId

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *SysVersion) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(sysVersionWorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
