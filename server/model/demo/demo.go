package demo

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type Demo struct {
	common.BaseModel
	Name string `json:"name" gorm:"index;not null;comment:demo名称"`
}

func (d *Demo) TableName() string {
	return "demo"
}

var demoWorkerId int64 = 100

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *Demo) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(demoWorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
