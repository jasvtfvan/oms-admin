package request

// 修改密码的结构 { oldPassword: "", newPassword: "", reNewPassword: "" }
type PwdSecret struct {
	Secret string `json:"secret" binding:"required"`
}

// 比较2个secret，每个都是由{"username":"xxx","password":"xxxxxx"}用rsa加密的
type CompareSecret struct {
	S1 string `json:"s1" binding:"required"`
	S2 string `json:"s2" binding:"required"`
}

// 用于登录自动绑定参数
type Login struct {
	// 由用户名和密码组成，例如: {"username":"xxx","password":"xxxxxx"}
	Secret    string `json:"secret" binding:"required"`
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// 用于登录时解析用户信息
type UserCredentials struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// 重置用户密码的结构
type ResetUserPassword struct {
	ID       int    `json:"id" binding:"required"` // id
	Password string `json:"password,omitempty"`    // 密码，omitempty该字段为空时在序列化时忽略它
}
