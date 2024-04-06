package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type OperationRecordDao struct{}

func (*OperationRecordDao) CreateOperationRecord(sysOperationRecord system.SysOperationRecord) error {
	db := global.OMS_DB
	return db.Create(&sysOperationRecord).Error
}
