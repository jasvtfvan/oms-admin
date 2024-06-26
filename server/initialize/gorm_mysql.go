package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	env := global.OMS_CONFIG.System.Env
	m := global.OMS_CONFIG.Mysql
	if m.DebugDbName == "" {
		m.DebugDbName = "oms"
	}
	if m.DbName == "" {
		m.DbName = "oms"
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(env), // DSN data source name
		DefaultStringSize:         191,        // string 类型字段的默认长度
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		// 连接数据库失败，写入fatal日志，因为代码错误，会导致程序不能正确运行
		global.OMS_LOG.Fatal("Gorm数据库连接失败")
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
