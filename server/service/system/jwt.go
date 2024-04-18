package system

import (
	sysModel "github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/utils"
	jwtRedis "github.com/jasvtfvan/oms-admin/server/utils/redis/jwt"
)

type JWTService interface {
	GenerateToken(sysUser *sysModel.SysUser) (string, error)
	DelStore(username string) error
}

type JWTServiceImpl struct{}

func (*JWTServiceImpl) DelStore(username string) error {
	var jwtStore = jwtRedis.GetRedisStore()
	return jwtStore.Del(username)
}

// sysUser 使用结构体对象，会创建一个副本；使用指针，会复用对象，只创建一个指针变量
func (*JWTServiceImpl) GenerateToken(sysUser *sysModel.SysUser) (string, error) {
	var jwtStore = jwtRedis.GetRedisStore()
	var sysGroupCodes []string
	for _, v := range sysUser.SysGroups {
		sysGroupCodes = append(sysGroupCodes, v.OrgCode)
	}
	var sysRoleCodes []string
	for _, v := range sysUser.SysRoles {
		sysRoleCodes = append(sysRoleCodes, v.RoleCode)
	}

	j := utils.NewJWT()
	claims := j.CreateClaims(utils.BaseClaims{
		ID:           sysUser.ID,
		Username:     sysUser.Username,
		LogOperation: sysUser.LogOperation,
		Groups:       sysGroupCodes,
		Roles:        sysRoleCodes,
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
