package system

import (
	sysModel "github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

type JWTService interface {
	GenerateToken(sysUser *sysModel.SysUser) (string, error)
}

type JWTServiceImpl struct{}

// sysUser 使用结构体对象，会创建一个副本；使用指针，会服用对象，只创建一个指针变量
func (*JWTServiceImpl) GenerateToken(sysUser *sysModel.SysUser) (string, error) {
	j := utils.NewJWT()
	claims := j.CreateClaims(utils.BaseClaims{
		ID:       sysUser.ID,
		Username: sysUser.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}
	err = jwtStore.Set(sysUser.Username, token)
	if err != nil {
		return "", err
	}
	return token, nil
}
