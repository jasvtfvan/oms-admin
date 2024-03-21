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

var demoWorkerId int64 = 0

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *Demo) BeforeCreate(db *gorm.DB) error {
	snowflakeWorker := utils.NewSnowflakeWorker(demoWorkerId)
	s.BaseModel.ID = uint(snowflakeWorker.NextId())
	return nil
}
