package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type OperationRecordService interface {
	CreateSysOperationRecord(system.SysOperationRecord) error
}

type OperationRecordServiceImpl struct{}

func (*OperationRecordServiceImpl) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) error {
	return operationRecordDao.CreateOperationRecord(sysOperationRecord)
}
