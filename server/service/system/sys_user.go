package system

import (
	"errors"

	sysDao "github.com/jasvtfvan/oms-admin/server/dao/system"
	sysModel "github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/utils"
	jwtRedis "github.com/jasvtfvan/oms-admin/server/utils/redis/jwt"
)

type UserService interface {
	Login(username string, password string) (*sysModel.SysUser, error)
	DeleteUser(id uint) error
	DisableUser(id uint) error
}

type UserServiceImpl struct{}

func (*UserServiceImpl) DisableUser(id uint) error {
	row, err := sysDao.DeleteUser(id)
	if err != nil {
		return err
	}
	if row != 0 {
		return errors.New("没有对应数据")
	}
	sysUser, err := sysDao.FindUserById(id)
	if err != nil {
		return err
	}
	if sysUser == nil {
		return errors.New("没有对应用户数据")
	}
	jwtStore := jwtRedis.GetRedisStore()
	err = jwtStore.Del(sysUser.Username)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceImpl) DeleteUser(id uint) error {
	row, err := sysDao.DeleteUser(id)
	if err != nil {
		return err
	}
	if row != 0 {
		return errors.New("没有对应数据")
	}
	sysUser, err := sysDao.FindUserById(id)
	if err != nil {
		return err
	}
	if sysUser == nil {
		return errors.New("没有对应用户数据")
	}
	jwtStore := jwtRedis.GetRedisStore()
	err = jwtStore.Del(sysUser.Username)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceImpl) Login(username string, password string) (*sysModel.SysUser, error) {
	var sysUser *sysModel.SysUser
	sysUser, err := sysDao.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(password, sysUser.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return sysUser, err
}
