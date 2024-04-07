package response

type SysCaptcha struct {
	CaptchaId     string `json:"captchaId"`     // 验证码id
	PicPath       string `json:"picPath"`       // 验证码图片
	CaptchaLength int    `json:"captchaLength"` // 验证码长度
	OpenCaptcha   bool   `json:"openCaptcha"`   // 是否开启验证码验证
}
