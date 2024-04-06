package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type UserRoleDao struct{}

func (*UserRoleDao) FindUserRolesByUserId(userId uint) ([]system.SysUserRole, error) {
	db := global.OMS_DB
	var sysUserRoles []system.SysUserRole
	err := db.Where("sys_user_id = ?", userId).Find(&sysUserRoles).Error
	if err != nil {
		return nil, err
	}
	return sysUserRoles, nil
}
