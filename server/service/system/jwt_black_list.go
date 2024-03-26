package system

import (
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"go.uber.org/zap"
)

type JWTService interface {
	JsonInBlackList(JWTList system.JWTBlackList) error
}

type JWTServiceImpl struct{}

func (*JWTServiceImpl) JsonInBlackList(JWTList system.JWTBlackList) (err error) {
	db := global.OMS_DB
	err = db.Create(&JWTList).Error

	return err
}

func LoadAll() {
	var data []string
	db := global.OMS_DB
	err := db.Model(&system.JWTBlackList{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.OMS_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		// global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
