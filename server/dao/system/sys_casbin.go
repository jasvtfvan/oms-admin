package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"gorm.io/gorm"
)

type CasbinDao struct{}

func (*CasbinDao) BatchDelete(tx *gorm.DB, roleCode string, groupCode string) error {
	return tx.Where("ptype = 'p' and v0 = ? and v1 = ?", roleCode, groupCode).Delete(&system.SysCasbin{}).Error
}

func (*CasbinDao) BatchInsert(tx *gorm.DB, casbinRules []system.SysCasbin) error {
	return tx.Create(&casbinRules).Error
}
