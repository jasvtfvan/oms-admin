package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`             // 环境值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseTls       bool   `mapstructure:"use-tls" json:"use-tls" yaml:"use-tls"`    // 开启https，使用tls证书
	TlsCert      string `mapstructure:"tls-cert" json:"tls-cert" yaml:"tls-cert"` // crt文件
	TlsKey       string `mapstructure:"tls-key" json:"tls-key" yaml:"tls-key"`    // key文件
}
