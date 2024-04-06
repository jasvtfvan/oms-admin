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
	InitOrderLevel3 = 100000
)

// 升级顺序 ！注意：升级代码-结构/数据 <==保持同步==> 初始化代码-结构/数据，
// 通过同步，部署其他企业时，一次性初始化，无需再执行升级程序
const (
	UpdateOrderLevel1 = 10
	UpdateOrderLevel2 = 1000
	UpdateOrderLevel3 = 100000
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
	InitOrderSysGroup     = InitOrderLevel1 + 2
	InitOrderSysRole      = InitOrderLevel1 + 3
	InitOrderSysUser      = InitOrderLevel1 + 4
	InitOrderSysUserGroup = InitOrderLevel1 + 5
	InitOrderSysUserRole  = InitOrderLevel1 + 6
	InitOrderSysCasbin    = InitOrderLevel1 + 7
)
const (
	InitOrderRegisterTables = InitOrderLevel2 + 1
)
const (
	InitOrderDemo = InitOrderLevel3 + 1
)

// 升级顺序定义
const (
	UpdateOrderSysVersion = UpdateOrderLevel1 + 1
)
const (
	UpdateOrderRegisterTables = UpdateOrderLevel2 + 1
)
const (
	UpdateOrderDemo = UpdateOrderLevel3 + 1
)

// ID WorkerId
const (
	SysVersionWorkerId = WorkerIdLevel1 + 1
	SysGroupWorkerId   = WorkerIdLevel1 + 2
	SysRoleWorkerId    = WorkerIdLevel1 + 3
	SysUserWorkerId    = WorkerIdLevel1 + 4
	SysCasbinWorkerId  = WorkerIdLevel1 + 5
)
const (
	DemoWorkerId = WorkerIdLevel2 + 1
)
