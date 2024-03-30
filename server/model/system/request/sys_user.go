package request

type Login struct {
	// 由用户名和密码组成，例如: {"username":"oms_admin","password":"Oms123Admin456"}
	Secret    string `json:"secret"`
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
