package system

import (
	sysDao "github.com/jasvtfvan/oms-admin/server/dao/system"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type OperationRecordService interface {
	CreateSysOperationRecord(system.SysOperationRecord) error
}

type OperationRecordServiceImpl struct{}

func (*OperationRecordServiceImpl) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) error {
	return sysDao.CreateOperationRecord(sysOperationRecord)
}
