package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/utils/freecache"
)

var freecacheStore = freecache.GetStoreDefault()

type CacheApi struct{}

// DoTestCache
// @Tags	test
// @Summary	测试local_cache
// @Produce	application/json
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}	"返回结果信息"
// @Router	/cache/test-cache [post]
func (*CacheApi) DoTestCache(c *gin.Context) {
	/*
		bool测试
	*/
	var bool1 freecache.Bool = true
	freecacheStore.Set("boolKey", bool1)
	var boolStruct1 freecache.Bool
	boolOutput1 := freecacheStore.Get("boolKey", &boolStruct1).(*freecache.Bool)
	if !freecacheStore.Verify("boolKey", &bool1, boolOutput1) {
		response.Fail(nil, "bool未通过测试", c)
		return
	}
	var boolStruct2 freecache.Bool
	boolOutput2 := freecacheStore.Get("boolKey", boolStruct2)
	if !freecacheStore.Verify("boolKey", true, boolOutput2) {
		response.Fail(nil, "bool未通过测试", c)
		return
	}

	/*
		string测试
	*/
	var string1 freecache.String = "hello world"
	freecacheStore.Set("stringKey", string1)
	var stringStruct1 freecache.String
	stringOutput1 := freecacheStore.Get("stringKey", &stringStruct1).(*freecache.String)
	if !freecacheStore.Verify("stringKey", &string1, stringOutput1) {
		response.Fail(nil, "string未通过测试", c)
		return
	}
	var stringStruct2 freecache.String
	stringOutput2 := freecacheStore.Get("stringKey", stringStruct2)
	if !freecacheStore.Verify("stringKey", "hello world", stringOutput2) {
		response.Fail(nil, "string未通过测试", c)
		return
	}

	/*
		结构体测试
	*/
	type Student struct {
		Name string `json:"name"`
	}
	s1 := Student{
		Name: "张三",
	}
	freecacheStore.Set("zs", s1)
	var sStruct Student
	// 需要把结构体的地址传入，否则输出结构会自动转成map
	sOutput := freecacheStore.Get("zs", &sStruct).(*Student)
	// fmt.Println(fmt.Sprintf("%+v:%+v", sOutput, *sOutput))
	if !freecacheStore.Verify("zs", &s1, sOutput) {
		response.Fail(nil, "结构体未通过测试", c)
		return
	}

	/*
		map测试
	*/
	map1 := make(map[string]string)
	map1["a"] = "hello"
	map1["b"] = "world"
	freecacheStore.Set("map1", map1)
	map2 := make(map[string]string)
	mOutput := freecacheStore.Get("map1", &map2).(*map[string]string)
	// fmt.Println(fmt.Sprintf("%+v:%+v", mOutput, *mOutput))
	if !freecacheStore.Verify("map1", &map1, mOutput) {
		response.Fail(nil, "map未通过测试", c)
		return
	}

	response.Success(nil, "全体通过测试", c)
}
