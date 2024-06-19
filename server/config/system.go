package config

type System struct {
	AuthCache    string `mapstructure:"auth-cache" json:"auth-cache" yaml:"auth-cache"` // token缓存的位置
	Username     string `mapstructure:"username" json:"username" yaml:"username"`       // 超级管理员账号
	Password     string `mapstructure:"password" json:"password" yaml:"password"`       // 超级管理员密码
	InitPwd      string `mapstructure:"init-pwd" json:"init-pwd" yaml:"init-pwd"`       // 初始化系统的密码init/db接口
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`          // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseTls       bool   `mapstructure:"use-tls" json:"use-tls" yaml:"use-tls"`    // 开启https，使用tls证书
	TlsCert      string `mapstructure:"tls-cert" json:"tls-cert" yaml:"tls-cert"` // crt文件
	TlsKey       string `mapstructure:"tls-key" json:"tls-key" yaml:"tls-key"`    // key文件
}
