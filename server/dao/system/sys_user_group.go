package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type UserGroupDao struct{}

func (*UserGroupDao) FindUserGroupsByUserID(userId uint) ([]system.SysUserGroup, error) {
	db := global.OMS_DB
	var sysUserGroups []system.SysUserGroup
	err := db.Where("sys_user_id = ?", userId).Find(&sysUserGroups).Error
	if err != nil {
		return nil, err
	}
	return sysUserGroups, nil
}
