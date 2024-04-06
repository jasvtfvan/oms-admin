package system

import (
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type SysOperationRecord struct {
	common.BaseModel
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	ErrorMessage string        `json:"errorMessage" form:"errorMessage" gorm:"column:error_message;comment:错误信息"`    // 错误信息
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                 // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                 // 响应Body
	UserID       int           `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id"`                      // 用户id
	User         SysUser       `json:"user"`                                                                         // gorm belongs to
}

func (s *SysOperationRecord) TableName() string {
	return "sys_operation_record"
}

var sysOperationRecordWorkId int64 = global.SysOperationRecordWorkId

// BeforeCreate 钩子，在创建记录前设置自定义的ID
func (s *SysOperationRecord) BeforeCreate(db *gorm.DB) error {
	if s.ID == 0 {
		snowflakeWorker := utils.NewSnowflakeWorker(sysOperationRecordWorkId)
		s.ID = uint(snowflakeWorker.NextId())
	}
	return nil
}
