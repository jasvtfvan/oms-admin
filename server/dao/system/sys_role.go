package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type RoleDao struct{}

func (*RoleDao) FindRolesByCodes(sysRoleCodes []string) ([]system.SysRole, error) {
	db := global.OMS_DB
	var sysRoles []system.SysRole
	err := db.Where("enable = true and role_code in ?", sysRoleCodes).Order("sort").Find(&sysRoles).Error
	if err != nil {
		return nil, err
	}
	return sysRoles, nil
}

func (*RoleDao) FindRolesByIds(sysRoleIds []uint) ([]system.SysRole, error) {
	db := global.OMS_DB
	var sysRoles []system.SysRole
	err := db.Where("enable = true and id in ?", sysRoleIds).Order("sort").Find(&sysRoles).Error
	if err != nil {
		return nil, err
	}
	return sysRoles, nil
}
