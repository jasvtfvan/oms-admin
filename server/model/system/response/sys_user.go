package response

import "github.com/jasvtfvan/oms-admin/server/model/system"

type Login struct {
	User  system.SysUser `json:"user"`
	Token string         `json:"token"`
}
