package system

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	sysReq "github.com/jasvtfvan/oms-admin/server/model/system/request"
	sysRes "github.com/jasvtfvan/oms-admin/server/model/system/response"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"github.com/jasvtfvan/oms-admin/server/utils/crypto"
	"github.com/jasvtfvan/oms-admin/server/utils/redis/captcha"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var captchaBuildBase64Store = captcha.GetBuildBase64Store()
var captchaLoginCountStore = captcha.GetLoginCountStore()
var captchaBuildCountStore = captcha.GetBuildCountStore()

type UserApi struct{}

// ResetPassword
// @Tags	user
// @Summary	重置密码
// @Produce	application/json
// @Param	data	body	sysReq.ResetUserPassword	true	"id（必填），password（必填）"
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}	"返回加密的密码，前端自行解密"
// @Router	/user/reset-pwd [put]
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
	encryptedPassword, username, err := userService.ResetPassword(uint(req.ID), newPassword)
	if err != nil {
		global.OMS_LOG.Error("重置用户密码失败:"+username, zap.Error(err))
		response.Fail(gin.H{encryptedPassword: encryptedPassword}, "操作失败:"+err.Error(), c)
		return
	}
	if err = jwtService.DelStore(username); err != nil {
		global.OMS_LOG.Error("删除jwt缓存失败:"+username, zap.Error(err))
	}
	response.Success(gin.H{encryptedPassword: encryptedPassword}, "操作成功", c)
}

// EnableUser
// @Tags	user
// @Summary	启用用户
// @Produce	application/json
// @Param	id	path	int	true	"用户ID"
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}
// @Router	/user/enable/{id} [put]
func (u *UserApi) EnableUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(nil, "id不能为空", c)
		return
	}
	idInt, _ := strconv.Atoi(id)
	username, err := userService.EnableUser(uint(idInt))
	if err != nil {
		global.OMS_LOG.Error("启用用户失败:"+username, zap.Error(err))
		response.Fail(nil, "操作失败:"+err.Error(), c)
		return
	}
	response.Success(nil, "操作成功", c)
}

// DisableUser
// @Tags	user
// @Summary	禁用用户
// @Produce	application/json
// @Param	id	path	int	true	"用户ID"
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}
// @Router	/user/disable/{id} [put]
func (u *UserApi) DisableUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(nil, "id不能为空", c)
		return
	}
	idInt, _ := strconv.Atoi(id)
	username, err := userService.DisableUser(uint(idInt))
	if err != nil {
		global.OMS_LOG.Error("禁用用户失败:"+username, zap.Error(err))
		response.Fail(nil, "操作失败:"+err.Error(), c)
		return
	}
	if err = jwtService.DelStore(username); err != nil {
		global.OMS_LOG.Error("删除jwt缓存失败:"+username, zap.Error(err))
	}
	response.Success(nil, "操作成功", c)
}

// DisableUser
// @Tags	user
// @Summary	删除用户
// @Produce	application/json
// @Param	id	path	int	true	"用户ID"
// @Success	200	{object}	response.Response{code=int,data=any,msg=string}
// @Router	/user/disable/{id} [delete]
func (u *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(nil, "id不能为空", c)
		return
	}
	idInt, _ := strconv.Atoi(id)
	username, err := userService.DeleteUser(uint(idInt))
	if err != nil {
		global.OMS_LOG.Error("删除用户失败:"+username, zap.Error(err))
		response.Fail(nil, "操作失败:"+err.Error(), c)
		return
	}
	if err = jwtService.DelStore(username); err != nil {
		global.OMS_LOG.Error("删除jwt缓存失败:"+username, zap.Error(err))
	}
	response.Success(nil, "操作成功", c)
}

// Captcha
// @Tags	base
// @Summary	重置密码
// @Produce	application/json
// @Success	200	{object}	response.Response{code=int,data=sysRes.SysCaptcha,msg=string}	"返回验证码信息"
// @Router	/base/captcha [post]
func (u *UserApi) Captcha(c *gin.Context) {
	key := c.ClientIP()                                                       // 使用ip当验证码的key
	openCaptchaBuildCountMax := global.OMS_CONFIG.Captcha.OpenCaptchaBuildMax // 验证码次数最多可以生成多少次，超过后锁定timeout时长
	captchaBuildCountStore = captchaBuildCountStore.UseWithCtx(c)
	buildCountCount := captchaBuildCountStore.GetCount(key) // 验证码次数
	if buildCountCount <= 0 {
		captchaBuildCountStore.InitCount(key)
	}

	if buildCountCount >= openCaptchaBuildCountMax {
		global.OMS_LOG.Error("获取验证码，太频繁", zap.Error(errors.New("ip: "+key)))
		response.Fail(nil, "操作太频繁，过一阵再来尝试。", c)
		return
	}

	openCaptcha := global.OMS_CONFIG.Captcha.OpenCaptcha // 防爆次数
	captchaLoginCountStore = captchaLoginCountStore.UseWithCtx(c)
	count := captchaLoginCountStore.GetCount(key) // 验证码次数
	if count <= 0 {
		captchaLoginCountStore.InitCount(key)
	}
	var isOpen bool // 为0直接开启防爆 或者 如果超过防爆次数，则开启防爆
	// 如果当前是第4次，count上次计算为3，这次需要>=
	if openCaptcha == 0 || count >= openCaptcha {
		isOpen = true
	}

	height := global.OMS_CONFIG.Captcha.ImgHeight                                      // 验证码高度
	width := global.OMS_CONFIG.Captcha.ImgWidth                                        // 验证码宽度
	keyLong := global.OMS_CONFIG.Captcha.KeyLong                                       // 验证码长度
	driver := base64Captcha.NewDriverDigit(height, width, keyLong, 0.7, 80)            // 初始化driver
	captcha := base64Captcha.NewCaptcha(driver, captchaBuildBase64Store.UseWithCtx(c)) // 新建验证码对象
	id, b64s, _, err := captcha.Generate()                                             // 生成验证码信息
	if err != nil {
		captchaBuildCountStore.AddCount(key) // 生成验证码次数+1
		global.OMS_LOG.Error("验证码获取失败:", zap.Error(err))
		response.Fail(nil, "验证码获取失败", c)
		return
	}

	captchaBuildCountStore.AddCount(key) // 生成验证码次数+1
	response.Success(sysRes.SysCaptcha{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.OMS_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   isOpen,
	}, "验证码获取成功", c)
}

// Login
// @Tags	base
// @Summary	用户登录
// @Produce	application/json
// @Param	data	body	sysReq.Login	true	"secret（必填），验证码+验证码id（选填）"
// @Success	200	{object}	response.Response{code=int,data=sysRes.Login,msg=string}	"返回用户信息,token"
// @Router	/base/login [post]
func (u *UserApi) Login(c *gin.Context) {
	key := c.ClientIP()                                        // 使用ip当验证码的key
	openCaptcha := global.OMS_CONFIG.Captcha.OpenCaptcha       // 防爆次数
	openCaptchaMax := global.OMS_CONFIG.Captcha.OpenCaptchaMax // 最大次数，超过后锁定timeout时长
	captchaLoginCountStore = captchaLoginCountStore.UseWithCtx(c)
	count := captchaLoginCountStore.GetCount(key) // 验证码次数
	if count <= 0 {
		captchaLoginCountStore.InitCount(key)
	}

	if count >= openCaptchaMax { // 超过最大次数，锁定
		global.OMS_LOG.Error("错误登录，太频繁", zap.Error(errors.New("ip: "+key)))
		response.Fail(nil, "错误太过频繁，过一阵再来尝试。", c)
		return
	}

	var req sysReq.Login
	err := c.ShouldBindJSON(&req) // 自动绑定
	if err != nil {
		captchaLoginCountStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.LoginVerify)
	if err != nil {
		captchaLoginCountStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, err.Error(), c)
		return
	}

	decrypted := crypto.RsaDecrypt(req.Secret)
	var userCreds sysReq.UserCredentials
	err = json.Unmarshal([]byte(decrypted), &userCreds)
	if err != nil {
		captchaLoginCountStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "参数格式错误: "+err.Error(), c)
		return
	}

	if userCreds.Username == "" || userCreds.Password == "" {
		captchaLoginCountStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "用户名或密码错误", c)
		return
	}

	var isOpen bool // 为0直接开启防爆 或者 如果超过防爆次数，则开启防爆
	// 如果当前是第4次，count上次计算为3，这次需要>=
	if openCaptcha == 0 || count >= openCaptcha {
		isOpen = true
	}

	// 开启后，验证码信息不能为空
	if isOpen && (req.CaptchaId == "" || req.Captcha == "") {
		captchaLoginCountStore.AddCount(key) // 验证码次数+1
		response.Fail(nil, "请输入验证码", c)
		return
	}

	// 如果防爆尚未开启，直接进行登录；如果防爆开启，则需要验证码验证，验证码只能使用一次
	if !isOpen || captchaBuildBase64Store.Verify(req.CaptchaId, req.Captcha, true) {
		// 登录service，失败或用户禁用，则返回错误信息，验证码次数+1
		user, err := userService.Login(userCreds.Username, userCreds.Password)
		if err != nil {
			global.OMS_LOG.Error("登录失败，用户名或密码错误", zap.Error(err))
			captchaLoginCountStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "用户名或密码错误", c)
			return
		}
		if !user.Enable {
			global.OMS_LOG.Error("登录失败，用户被禁用", zap.Error(err))
			captchaLoginCountStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "用户被禁用", c)
			return
		}
		// 获取用户所有群组
		user.SysGroups, err = groupService.FindGroupsByUserID(user.ID)
		if err != nil {
			global.OMS_LOG.Error("登录失败，用户组织查询失败", zap.Error(err))
			captchaLoginCountStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "用户组织查询失败", c)
			return
		}
		// 获取用户所有角色
		user.SysRoles, err = roleService.FindRolesByUserID(user.ID)
		if err != nil {
			global.OMS_LOG.Error("登录失败，用户角色查询失败", zap.Error(err))
			captchaLoginCountStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "用户角色查询失败", c)
			return
		}
		// 登录成功，清除验证码次数，创建token，返回正确信息
		token, err := jwtService.GenerateToken(user)
		if err != nil {
			global.OMS_LOG.Error("获取token失败", zap.Error(err))
			captchaLoginCountStore.AddCount(key) // 验证码次数+1
			response.Fail(nil, "令牌获取失败", c)
			return
		}
		global.OMS_LOG.Info("登录成功")
		captchaBuildCountStore.DelCount(key) // 清除生成次数
		captchaLoginCountStore.DelCount(key) // 清除验证码次数
		// 将group和role返回
		loginGroups := []sysRes.LoginGroups{}
		for _, grp := range user.SysGroups { // 迭代group
			loginRoles := []sysRes.LoginRole{}
			for _, rl := range user.SysRoles { // 迭代role
				if rl.SysGroupID == grp.ID {
					loginRoles = append(loginRoles, sysRes.LoginRole{
						RoleName: rl.RoleName,
						RoleCode: rl.RoleCode,
						Sort:     rl.Sort,
					})
				}
			}
			loginGroups = append(loginGroups, sysRes.LoginGroups{
				ShortName: grp.ShortName,
				OrgCode:   grp.OrgCode,
				Sort:      grp.Sort,
				SysRoles:  loginRoles,
			})
		}
		response.Success(sysRes.Login{
			User: sysRes.LoginUser{
				Username:     user.Username,
				NickName:     user.NickName,
				Avatar:       user.Avatar,
				Phone:        user.Phone,
				Email:        user.Email,
				IsAdmin:      user.IsAdmin,
				LogOperation: user.LogOperation,
				SysGroups:    loginGroups,
			},
			Token: token,
		}, "登录成功", c)
		return
	}

	captchaLoginCountStore.AddCount(key) // 验证码次数+1
	response.Fail(nil, "验证码错误，请重新获取", c)
}
