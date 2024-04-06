package initializer

import (
	"strings"

	"github.com/jasvtfvan/oms-admin/server/global"
)

func GetRootRoleCode() string {
	return GetRootGroupCode() + "_admin"
}

func GetRootGroupCode() string {
	rootUsername := GetRootUsername()
	var OrgCode = "root"
	// 如果系统管理员名字以_admin结尾则以_admin前边为根组织的编号
	if len(rootUsername) > 6 && strings.HasSuffix(rootUsername, "_admin") {
		OrgCode = strings.TrimSuffix(rootUsername, "_admin")
	}
	return OrgCode
}

func GetRootUsername() string {
	return global.OMS_CONFIG.System.Username
}
