package global

/*
	根据给定的顺序，链式传递，比如：A调用InitOrderLevel1+1，则B=A+1，C=B+1，D=C+1以此类推，避免顺序重复
*/

/*
	顺序起始参数
*/
// 初始化顺序
const (
	InitOrderLevel1 = 10
	InitOrderLevel2 = 1000
)

// 升级顺序
const (
	UpdateOrderLevel1 = 10
	UpdateOrderLevel2 = 1000
)

// ID WorkerId
const (
	WorkerIdLevel1 int64 = 0
	WorkerIdLevel2 int64 = 100
)

/*
	==============================
	顺序衍生参数，用于各个业务中
*/
// 初始化顺序定义
const (
	InitOrderSysVersion   = InitOrderLevel1 + 1
	InitOrderJWTBlackList = InitOrderSysVersion + 1
	InitOrderSysGroup     = InitOrderJWTBlackList + 1
	InitOrderSysRole      = InitOrderSysGroup + 1
	InitOrderSysUser      = InitOrderSysRole + 1
	InitOrderSysUserGroup = InitOrderSysUser + 1
	InitOrderSysUserRole  = InitOrderSysUserGroup + 1
)
const (
	InitOrderDemo = InitOrderLevel2 + 1
)

// 升级顺序定义
const (
	UpdateOrderSysVersion = UpdateOrderLevel1 + 1
)
const (
	UpdateOrderDemo = UpdateOrderLevel2 + 1
)

// ID WorkerId
const (
	SysVersionWorkerId      = WorkerIdLevel1 + 1
	SysJWTBlacklistWorkerId = SysVersionWorkerId + 1
	SysGroupWorkerId        = SysJWTBlacklistWorkerId + 1
	SysRoleWorkerId         = SysGroupWorkerId + 1
	SysUserWorkerId         = SysRoleWorkerId + 1
)
const (
	DemoWorkerId = WorkerIdLevel2 + 1
)
