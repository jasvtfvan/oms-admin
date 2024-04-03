package dao

import "github.com/jasvtfvan/oms-admin/server/dao/system"

type DaoGroup struct {
	System system.ServiceGroup
}

var DaoGroupApp = new(DaoGroup)
