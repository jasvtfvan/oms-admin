package global

/*
	根据给定的顺序，链式传递，比如：A调用InitOrderSystem，则B=A+1，C=B+1，D=C+1以此类推，避免顺序重复
*/

/*
	顺序起始参数
*/
// 初始化顺序
const (
	InitOrderSystem = 10
	InitOrderDemo   = 1000
)

// 升级顺序
const (
	UpdateOrderSystem = 10
	UpdateOrderDemo   = 1000
)

// ID WorkerId
const (
	SystemWorkerId int64 = 0
	DemoWorkerId   int64 = 100
)

/*
	==============================
	顺序衍生参数，用于各个业务中
*/
// 初始化顺序定义
const (
	InitOrderSysVersion   = InitOrderSystem + 1
	InitOrderJWTBlackList = InitOrderSysVersion + 1
	InitOrderSysGroup     = InitOrderJWTBlackList + 1
	InitOrderSysRole      = InitOrderSysGroup + 1
	InitOrderSysUser      = InitOrderSysRole + 1
	InitOrderSysUserGroup = InitOrderSysUser + 1
	InitOrderSysUserRole  = InitOrderSysUserGroup + 1
)
