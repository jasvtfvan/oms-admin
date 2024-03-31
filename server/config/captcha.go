package config

type Captcha struct {
	KeyLong             int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                           // 验证码长度
	ImgWidth            int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                        // 验证码宽度
	ImgHeight           int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                                     // 验证码高度
	OpenCaptcha         int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                               // 防爆破验证码开启次数，0代表每次登录都需要验证码，其他数字代表错误密码次数，如3代表错误三次后出现验证码
	OpenCaptchaMax      int `mapstructure:"open-captcha-max" json:"open-captcha-max" yaml:"open-captcha-max"`                   // 最多登录失败次数，超过则锁定ip时长为OpenCaptchaTimeout秒
	OpenCaptchaBuildMax int `mapstructure:"open-captcha-build-max" json:"open-captcha-build-max" yaml:"open-captcha-build-max"` // 最多生成验证码的次数，超过这个次数，还没有登录成功，则锁定ip时长为OpenCaptchaTimeout秒
	OpenCaptchaTimeout  int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"`       // 防爆破验证码超时时间，单位：s(秒)，超过这个时间，恢复直接登录
}
