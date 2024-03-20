package internal

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	systemService "github.com/jasvtfvan/oms-admin/server/service/system"
	"gorm.io/gorm"
)

func DataInserted[T any](ctx context.Context, instance T, query string, args ...string) bool {

	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		global.OMS_LOG.Fatal(systemService.ErrMissingDBContext.Error())
		return false
	}

	err := db.Where(query, args).First(instance).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func MigrateTable[T any](ctx context.Context, instance T) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(instance)
}

func TableCreated[T any](ctx context.Context, instance T) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		global.OMS_LOG.Fatal(systemService.ErrMissingDBContext.Error())
		return false
	}
	return db.Migrator().HasTable(instance)
}
