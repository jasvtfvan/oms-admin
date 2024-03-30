package system

// 该结构体不创建table，只有特殊api才需要授权
// 每次代码更新，技术人员可以根据具体情况，把需要授权的api，写到sys_api中
type SysCasbinApi struct {
	Path   string
	Method string
}
