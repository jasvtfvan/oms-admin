package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

type JWTService interface {
}

type JWTServiceImpl struct{}

func (*JWTServiceImpl) JsonInBlackList(JWTList system.JWTBlackList) (err error) {
	// db := global.OMS_DB

	return err
}
