package demo

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

type Demo1 struct {
	common.BaseModel
	Name string `json:"name" gorm:"index;not null;comment:demo名称"`
}

func (d *Demo1) TableName() string {
	return "demo1"
}

var demo1WorkerId int64 = demoWorkerId + 1

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *Demo1) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(demo1WorkerId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
