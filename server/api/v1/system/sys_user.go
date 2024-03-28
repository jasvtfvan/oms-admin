package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/model/system/request"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var user request.Login
	err := c.ShouldBindJSON(&user)
	// ip := c.ClientIP()

	if err != nil {
		response.Fail(nil, err.Error(), c)
		return
	}

}
