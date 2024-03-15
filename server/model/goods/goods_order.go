package goods

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
)

type GoodsOrder struct {
	common.BaseModel
	GoodsNum string `json:"goodsNum" gorm:"index;not null;comment:商品编码"`
	OrderNum string `json:"orderNum" gorm:"index;not null;comment:订单编码"`
}

func (GoodsOrder) TableName() string {
	return "goods_order"
}
