package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
)

type DemoApi struct{}

func (*DemoApi) Demo(c *gin.Context) {
	hello := demoService.Hello()
	response.Success(gin.H{hello: hello}, "这是一个Demo测试接口", c)
}
