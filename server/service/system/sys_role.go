package system

import (
	sysModel "github.com/jasvtfvan/oms-admin/server/model/system"
)

type RoleInstance struct {
	RoleService
}

var RoleServiceApp = &RoleInstance{
	RoleService: new(RoleServiceImpl),
}

type RoleService interface {
	FindRolesByUserID(uint) ([]sysModel.SysRole, error)
}

type RoleServiceImpl struct{}

func (*RoleServiceImpl) FindRolesByUserID(userId uint) ([]sysModel.SysRole, error) {
	sysUserRoles, err := userRoleDao.FindUserRolesByUserId(userId)
	if err != nil {
		return nil, err
	}

	var sysRoleIds []uint
	for _, v := range sysUserRoles {
		sysRoleIds = append(sysRoleIds, v.SysRoleID)
	}

	return roleDao.FindRolesByIds(sysRoleIds)
}
