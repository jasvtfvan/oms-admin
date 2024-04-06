package v1

import (
	"github.com/jasvtfvan/oms-admin/server/api/v1/system"
)

type ApiGroup struct {
	System system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
