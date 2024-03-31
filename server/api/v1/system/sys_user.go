package system

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	sysReq "github.com/jasvtfvan/oms-admin/server/model/system/request"
	sysRes "github.com/jasvtfvan/oms-admin/server/model/system/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"github.com/jasvtfvan/oms-admin/server/utils/redis/captcha"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var captchaStore = captcha.GetRedisStore()

type UserApi struct{}

func (u *UserApi) ResetPassword(c *gin.Context) {
	var req sysReq.ResetUserPassword
	err := c.ShouldBindJSON(&req) // 自动绑定
	if err != nil {
		response.Fail(nil, err.Error(), c)
		return
	}
	if req.ID == 0 {
		response.Fail(nil, "id不能为空", c)
		return
	}
	newPassword := req.Password
	encryptedPassword, err := userService.ResetPassword(uint(req.ID), newPassword)
	if err != nil {
		global.OMS_LOG.Error("重置用户密码失败id:"+strconv.Itoa(req.ID), zap.Error(err))
		response.Fail(gin.H{encryptedPassword: encryptedPassword}, "操作失败:"+err.Error(), c)
		return
	}
	response.Success(gin.H{encryptedPassword: encryptedPassword}, "操作成功", c)
}

func (u *UserApi) EnableUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(nil, "id不能为空", c)
		return
	}
	idInt, _ := strconv.Atoi(id)
	err := userService.EnableUser(uint(idInt))
	if err != nil {
		global.OMS_LOG.Error("启用用户失败id:"+id, zap.Error(err))
		response.Fail(nil, "操作失败:"+err.Error(), c)
		return
	}
	response.Success(nil, "操作成功", c)
}

func (u *UserApi) DisableUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(nil, "id不能为空", c)
		return
	}
	idInt, _ := strconv.Atoi(id)
	err := userService.DisableUser(uint(idInt))
	if err != nil {
		global.OMS_LOG.Error("禁用用户失败id:"+id, zap.Error(err))
		response.Fail(nil, "操作失败:"+err.Error(), c)
		return
	}
	response.Success(nil, "操作成功", c)
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(nil, "id不能为空", c)
		return
	}
	idInt, _ := strconv.Atoi(id)
	err := userService.DeleteUser(uint(idInt))
	if err != nil {
		global.OMS_LOG.Error("删除用户失败id:"+id, zap.Error(err))
		response.Fail(nil, "操作失败:"+err.Error(), c)
		return
	}
	response.Success(nil, "操作成功", c)
}

func (u *UserApi) Captcha(c *gin.Context) {
	openCaptcha := global.OMS_CONFIG.Captcha.OpenCaptcha // 防爆次数
	key := c.ClientIP()                                  // 使用ip当验证码的key
	count, ok := captchaStore.GetCount(key)              // 验证码次数
	if ok != nil {
		captchaStore.InitCount(key) // 初始化次数1
	}

	var isOpen bool // 为0直接开启防爆 或者 如果超过防爆次数，则开启防爆
	if openCaptcha == 0 || count > openCaptcha {
		isOpen = true
	}
	height := global.OMS_CONFIG.Captcha.ImgHeight                           // 验证码高度
	width := global.OMS_CONFIG.Captcha.ImgWidth                             // 验证码宽度
	keyLong := global.OMS_CONFIG.Captcha.KeyLong                            // 验证码长度
	driver := base64Captcha.NewDriverDigit(height, width, keyLong, 0.7, 80) // 初始化driver
	captcha := base64Captcha.NewCaptcha(driver, captchaStore.UseWithCtx(c)) // 新建验证码对象
	id, b64s, _, err := captcha.Generate()                                  // 生成验证码信息
	if err != nil {
		global.OMS_LOG.Error("验证码获取失败:", zap.Error(err))
		response.Fail(nil, "验证码获取失败", c)
		return
	}
	response.Success(sysRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.OMS_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   isOpen,
	}, "验证码获取成功", c)
}

func (u *UserApi) Login(c *gin.Context) {
	// 记录登录次数
	key := c.ClientIP()                                        // 使用ip当验证码的key
	openCaptcha := global.OMS_CONFIG.Captcha.OpenCaptcha       // 防爆次数
	openCaptchaMax := global.OMS_CONFIG.Captcha.OpenCaptchaMax // 最大次数，超过后锁定timeout时长
	count, ok := captchaStore.GetCount(key)                    // 验证码次数
	if ok != nil {
		captchaStore.InitCount(key) // 初始化次数1
	}

	if count > openCaptchaMax {
		captchaStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "错误太多频繁，过一阵再来尝试。也可以联系管理员", c)
		return
	}

	var req sysReq.Login
	err := c.ShouldBindJSON(&req) // 自动绑定
	if err != nil {
		captchaStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.LoginVerify)
	if err != nil {
		captchaStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, err.Error(), c)
		return
	}

	var userCreds sysReq.UserCredentials
	err = json.Unmarshal([]byte(req.Secret), &userCreds)
	if err != nil {
		captchaStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "参数格式错误: "+err.Error(), c)
		return
	}

	if userCreds.Username == "" || userCreds.Password == "" {
		captchaStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "用户名或密码错误", c)
		return
	}

	var isOpen bool // 为0直接开启防爆 或者 如果超过防爆次数，则开启防爆
	if openCaptcha == 0 || count > openCaptcha {
		isOpen = true
	}

	// 开启后，验证码信息不能为空
	if isOpen && (req.CaptchaId == "" || req.Captcha == "") {
		captchaStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "请输入验证码", c)
		return
	}

	// 如果防爆尚未开启，直接进行登录；如果防爆开启，则需要验证码验证，验证码只能使用一次
	if !isOpen || captchaStore.Verify(req.CaptchaId, req.Captcha, true) {
		// 登录service，失败或用户禁用，则返回错误信息，验证码次数+1
		user, err := userService.Login(userCreds.Username, userCreds.Password)
		if err != nil {
			global.OMS_LOG.Error("登录失败，用户名或密码错误", zap.Error(err))
			captchaStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "用户名或密码错误", c)
			return
		}
		if !user.Enable {
			global.OMS_LOG.Error("登录失败，用户被禁用", zap.Error(err))
			captchaStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "用户被禁用", c)
			return
		}
		// 登录成功，清除验证码次数，创建token，返回正确信息
		token, err := jwtService.GenerateToken(user)
		if err != nil {
			global.OMS_LOG.Error("获取token失败", zap.Error(err))
			captchaStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "令牌获取失败", c)
			return
		}
		global.OMS_LOG.Info("登录成功")
		captchaStore.DelCount(key) // 清除验证码次数
		response.Success(sysRes.Login{
			User:  *user,
			Token: token,
		}, "登录成功", c)
		return
	}

	captchaStore.AddCount(key) // 验证码次数+1
	response.Fail(nil, "验证码错误", c)
}
