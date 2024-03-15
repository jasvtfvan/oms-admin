package common

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"ID" gorm:"primarykey"`     // 主键ID
	CreatedAt time.Time      `gorm:"default:0"`                // 创建时间
	UpdatedAt time.Time      `gorm:"default:0"`                // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;default:0"` // 删除时间
}
