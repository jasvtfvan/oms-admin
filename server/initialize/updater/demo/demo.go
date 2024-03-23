package updater

import (
	"context"
	"errors"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/demo"
	"github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 更新顺序
const updateOrderDemo = global.UpdateOrderDemo + 1

type updateDemo struct{}

// UpdateData implements initialize.Updater.
func (u *updateDemo) UpdateData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

	result := db.Model(demo.Demo{}).Where("name = ?", "demo").
		Updates(demo.Demo{Desc: "更新的描述" + time.Now().Format(time.DateTime)})
	if err = result.Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + u.UpdaterName() + " 表数据更新失败")
	}

	next = context.WithValue(ctx, u.UpdaterName(), result.RowsAffected)
	return next, err
}

// UpdateTable implements initialize.Updater.
func (u *updateDemo) UpdateTable(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB
	return ctx, db.AutoMigrate(&demo.Demo{})
}

// UpdaterName implements initialize.Updater.
func (u *updateDemo) UpdaterName() string {
	return (&demo.Demo{}).TableName()
}

// auto run
func init() {
	initialize.RegisterUpdate(updateOrderDemo, &updateDemo{})
}
