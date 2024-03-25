package system

import (
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type JWTBlackList struct {
	ID        uint      `json:"ID" gorm:"primaryKey"` // 主键ID
	JWT       string    `json:"JWT" gorm:"type:text;comment:JWT"`
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

func (s *JWTBlackList) TableName() string {
	return "sys_jwt_blacklist"
}

var sysJWTBlacklistWorkerId int64 = global.SystemWorkerId

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *JWTBlackList) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(sysJWTBlacklistWorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
