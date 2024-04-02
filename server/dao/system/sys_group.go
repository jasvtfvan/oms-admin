package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type GroupDao struct{}

func FindGroupsByIds(sysGroupIds []uint) ([]system.SysGroup, error) {
	db := global.OMS_DB
	var sysGroups []system.SysGroup
	err := db.Where("enable = true and sys_group_id in ?", sysGroupIds).Order("sort").Find(&sysGroups).Error
	if err != nil {
		return nil, err
	}
	return sysGroups, nil
}
