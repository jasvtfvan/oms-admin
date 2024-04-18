package system

import (
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	sysModel "github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"github.com/jasvtfvan/oms-admin/server/utils/crypto"
)

type UserService interface {
	Login(username string, password string) (*sysModel.SysUser, error)
	DeleteUser(uint) (string, error)
	DisableUser(uint) (string, error)
	EnableUser(uint) (string, error)
	ResetPassword(uint, string) (string, string, error)
	FindUser(uint) (*sysModel.SysUser, error)
}

type UserServiceImpl struct{}

func (*UserServiceImpl) FindUser(id uint) (*sysModel.SysUser, error) {
	return userDao.FindUserById(id)
}

func (*UserServiceImpl) ResetPassword(id uint, newPassword string) (string, string, error) {
	sysUser, err := userDao.FindUserById(id)
	if err != nil {
		return "", "", err
	}
	if sysUser == nil {
		return "", "", errors.New("没有查到该用户")
	}
	if newPassword == "" {
		// 默认密码举例：zhangsan123456
		newPassword = sysUser.Username + "123456"
	}
	if len(newPassword) < 6 {
		return "", sysUser.Username, errors.New("密码不能低于6位")
	}
	password := utils.BcryptHash(newPassword)
	row, err := userDao.UpdatePassword(id, password)
	if err != nil {
		return "", sysUser.Username, err
	}
	if row != 0 {
		return "", sysUser.Username, errors.New("数据未响应")
	}
	encryptedPassword := crypto.AesEncrypt(newPassword)
	return encryptedPassword, sysUser.Username, nil
}

func (*UserServiceImpl) EnableUser(id uint) (string, error) {
	sysUser, err := userDao.FindUserById(id)
	if err != nil {
		return "", err
	}
	if sysUser == nil {
		return "", errors.New("没有查到该用户")
	}
	if global.OMS_CONFIG.System.Username == sysUser.Username {
		return sysUser.Username, errors.New("不能对系统管理员进行操作")
	}

	row, err := userDao.EnableUser(id)
	if err != nil {
		return sysUser.Username, err
	}
	if row == 0 {
		return sysUser.Username, errors.New("数据未响应")
	}
	return sysUser.Username, nil
}

func (*UserServiceImpl) DisableUser(id uint) (string, error) {
	sysUser, err := userDao.FindUserById(id)
	if err != nil {
		return "", err
	}
	if sysUser == nil {
		return "", errors.New("没有查到该用户")
	}
	if global.OMS_CONFIG.System.Username == sysUser.Username {
		return sysUser.Username, errors.New("不能对系统管理员进行操作")
	}

	row, err := userDao.DisableUser(id)
	if err != nil {
		return sysUser.Username, err
	}
	if row == 0 {
		return sysUser.Username, errors.New("数据未响应")
	}
	return sysUser.Username, nil
}

func (*UserServiceImpl) DeleteUser(id uint) (string, error) {
	sysUser, err := userDao.FindUserById(id)
	if err != nil {
		return "", err
	}
	if sysUser == nil {
		return "", errors.New("没有查到该用户")
	}
	if global.OMS_CONFIG.System.Username == sysUser.Username {
		return sysUser.Username, errors.New("不能对系统管理员进行操作")
	}

	row, err := userDao.DeleteUser(id)
	if err != nil {
		return sysUser.Username, err
	}
	if row == 0 {
		return sysUser.Username, errors.New("数据未响应")
	}
	return sysUser.Username, nil
}

func (*UserServiceImpl) Login(username string, password string) (*sysModel.SysUser, error) {
	var sysUser *sysModel.SysUser
	sysUser, err := userDao.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(password, sysUser.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return sysUser, err
}
