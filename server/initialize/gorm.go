package initialize

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/internal"
	"github.com/jasvtfvan/oms-admin/server/model/goods"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql() *gorm.DB {
	m := global.OMS_CONFIG.Mysql
	if m.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// 创建数据库(根据名称)
func CreateDatabase() {
	m := global.OMS_CONFIG.Mysql
	if m.DbName == "" {
		return
	}
	// 连接到MySQL服务器，注意这里并没有指定数据库名，因为我们只是想要执行创建数据库的命令
	db, err := sql.Open("mysql", m.EmptyDsn())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 确保连接是活跃的
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// 创建数据库的SQL语句
	createDbSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", m.DbName)

	// 执行SQL语句
	_, err = db.Exec(createDbSQL)
	if err != nil {
		panic(err)
	}

	fmt.Println(
		"Database oms has already exist or has been created successfully! " +
			"数据库'oms'已存在/创建成功!",
	)
}

func RegisterTables() {
	db := global.OMS_DB
	err := db.AutoMigrate(
		&system.SysUser{},
		&system.SysGroup{},
		&system.SysRole{},
		&goods.GoodsOrder{},
	)
	if err != nil {
		// global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	// global.GVA_LOG.Info("register table success")
}
