package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type GroupDao struct{}

func (*GroupDao) FindGroupsByCodes(sysGroupCodes []string) ([]system.SysGroup, error) {
	db := global.OMS_DB
	var sysGroups []system.SysGroup
	err := db.Where("enable = true and org_code in ?", sysGroupCodes).Order("sort").Find(&sysGroups).Error
	if err != nil {
		return nil, err
	}
	return sysGroups, nil
}

func (*GroupDao) FindGroupsByIds(sysGroupIds []uint) ([]system.SysGroup, error) {
	db := global.OMS_DB
	var sysGroups []system.SysGroup
	err := db.Where("enable = true and id in ?", sysGroupIds).Order("sort").Find(&sysGroups).Error
	if err != nil {
		return nil, err
	}
	return sysGroups, nil
}
