package system

import (
	sysModel "github.com/jasvtfvan/oms-admin/server/model/system"
)

type GroupService interface {
	FindGroupsByUserID(uint) ([]sysModel.SysGroup, error)
	FindGroupsByCodes([]string) ([]sysModel.SysGroup, error)
}

type GroupServiceImpl struct{}

func (*GroupServiceImpl) FindGroupsByCodes(sysGroupCodes []string) ([]sysModel.SysGroup, error) {
	return groupDao.FindGroupsByCodes(sysGroupCodes)
}

func (*GroupServiceImpl) FindGroupsByUserID(userId uint) ([]sysModel.SysGroup, error) {
	sysUserGroups, err := userGroupDao.FindUserGroupsByUserID(userId)
	if err != nil {
		return nil, err
	}

	var sysGroupIds []uint
	for _, v := range sysUserGroups {
		sysGroupIds = append(sysGroupIds, v.SysGroupID)
	}

	return groupDao.FindGroupsByIds(sysGroupIds)
}
