package initializer

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"gorm.io/gorm"
)

func DataInserted[T any](ctx context.Context, instance T, query string, args ...string) bool {
	db := global.OMS_DB

	var interfaceArgs []interface{}
	for _, arg := range args {
		interfaceArgs = append(interfaceArgs, arg)
	}
	err := db.Where(query, interfaceArgs...).First(instance).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func MigrateTable[T any](ctx context.Context, instance T) (next context.Context, err error) {
	db := global.OMS_DB
	return ctx, db.AutoMigrate(instance)
}

func TableCreated[T any](ctx context.Context, instance T) bool {
	db := global.OMS_DB
	return db.Migrator().HasTable(instance)
}
