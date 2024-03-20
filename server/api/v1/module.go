package v1

import (
	"github.com/jasvtfvan/oms-admin/server/api/v1/demo"
	"github.com/jasvtfvan/oms-admin/server/api/v1/system"
)

type ApiGroup struct {
	System system.ApiGroup
	Demo   demo.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
