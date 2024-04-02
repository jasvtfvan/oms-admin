package migrate

import (
	"github.com/jasvtfvan/oms-admin/server/model/demo"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

// 在initializer中没有注册，但需要初始化的空表，这些表不能删除
var InitMigrateTables = []interface{}{
	&system.SysOperationRecord{},
	&demo.Demo{},
}

// 在updater中没有注册，需要新增的空表；或者已经初始化，但只修改表结构不更新字段的表
// 这些表同时要复制到[InitMigrateTables]的尾部，但是不能覆盖已经写好的[InitMigrateTables]
var UpdateMigrateTables = []interface{}{
	&demo.Demo{},
}
