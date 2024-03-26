package initialize

import (
	"context"
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/global"
)

type MysqlUpdateHandler struct{}

// UpdateData implements TypedDbUpdateHandler.
func (m *MysqlUpdateHandler) UpdateData(ctx context.Context, updaters updaterSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, up := range updaters {
		if n, err := up.UpdateData(next); err != nil {
			// 此处不能panic，因为系统已经运行，此处记录日志，用于排查错误，修复后再次升级
			global.OMS_LOG.Error(fmt.Sprintf(UpdateDataFailed, UpdateMysql, up.UpdaterName(), err.Error()))
			return err
		} else {
			// 数据更新成功，写入info日志
			global.OMS_LOG.Info(fmt.Sprintf(UpdateDataSuccess, UpdateMysql, up.UpdaterName()))
			next = n
		}
	}
	return nil
}

// UpdateTables implements TypedDbUpdateHandler.
func (m *MysqlUpdateHandler) UpdateTables(ctx context.Context, updaters updaterSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, up := range updaters {
		if nt, err := up.UpdateTable(next); err != nil {
			global.OMS_LOG.Error(fmt.Sprintf(UpdateTableFailed, UpdateMysql, up.UpdaterName(), err.Error()))
			return err
		} else {
			next = nt
		}
	}
	return nil
}

func NewMysqlUpdateHandler() *MysqlUpdateHandler {
	return &MysqlUpdateHandler{}
}
