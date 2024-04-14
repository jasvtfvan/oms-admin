package request

type SysCaptcha struct {
	Width  int `json:"width"`  // 验证码图片宽度
	Height int `json:"height"` // 验证码图片高度
}
