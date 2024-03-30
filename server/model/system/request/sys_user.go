package request

// 用于登录自动绑定参数
type Login struct {
	// 由用户名和密码组成，例如: {"username":"oms_admin","password":"Oms123Admin456"}
	Secret    string `json:"secret"`
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// 用于登录时解析用户信息
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetUserPassword struct {
	ID       int    `json:"id"`
	Password string `json:"password,omitempty"` // omitempty该字段为空时在序列化时忽略它
}
