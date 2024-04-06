package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn(env string) string {
	if env == "release" {
		return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
	} else {
		return m.DebugUsername + ":" + m.DebugPassword + "@tcp(" + m.DebugPath + ":" + m.Port + ")/" + m.DebugDbName + "?" + m.Config
	}
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
