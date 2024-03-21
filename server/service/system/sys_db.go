package system

import (
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
)

type DBService interface {
	CheckDB() error
}

type DBServiceImpl struct{}

// 检查数据连接
func (*DBServiceImpl) CheckDB() (err error) {
	if global.OMS_DB == nil {
		return errors.New("DB为空，数据库未创建")
	} else {
		db, _ := global.OMS_DB.DB()
		if err = db.Ping(); err != nil {
			return err
		}
	}
	return err
}
