package config

type DsnProvider interface {
	Dsn() string
}

// GeneralDB 也被 Mysql 原样使用
type GeneralDB struct {
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	Config string `mapstructure:"config" json:"config" yaml:"config"` // 高级配置
	/* system.env == debug 使用开发数据库 */
	DebugDbName   string `mapstructure:"debug-db-name" json:"debug-db-name" yaml:"debug-db-name"`    // 数据库名
	DebugUsername string `mapstructure:"debug-username" json:"debug-username" yaml:"debug-username"` // 用户名
	DebugPassword string `mapstructure:"debug-password" json:"debug-password" yaml:"debug-password"` // 数据库密码
	DebugPath     string `mapstructure:"debug-path" json:"debug-path" yaml:"debug-path"`             // 数据库连接地址
	/* system.env == release 使用生产数据库 */
	DbName   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 数据库连接地址
	/* --------------------------------- */
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //数据库引擎，默认InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}
