package demo

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
)

type Demo struct {
	common.BaseModel
	Name string `json:"name" gorm:"index;not null;comment:demo名称"`
}

func (Demo) TableName() string {
	return "demo"
}
