package router

import "github.com/jasvtfvan/oms-admin/server/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
