package request

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

type ResetUserPassword struct {
	ID       int    `json:"id" binding:"required"` // id
	Password string `json:"password,omitempty"`    // 密码，omitempty该字段为空时在序列化时忽略它
}
