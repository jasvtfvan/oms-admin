package tests

import (
	"testing"

	"github.com/jasvtfvan/oms-admin/server/model/common/request"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

type PageInfoTest struct {
	PageInfo request.PageInfo
	Name     string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := utils.Rules{"Page": {utils.NotEmpty()}, "PageSize": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 0
	testInfo.PageInfo.PageSize = 0
	err := utils.Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉0值")
	}
	testInfo.Name = ""
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = utils.Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能正常检测name为空")
	}
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = utils.Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通过检测")
	}
}
