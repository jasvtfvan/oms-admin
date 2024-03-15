package core

import (
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/core/internal"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/jasvtfvan/oms-admin/server/global"
)

// Viper //
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper() *viper.Viper {
	var config string = internal.ConfigDefaultFile

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.OMS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.OMS_CONFIG); err != nil {
		panic(err)
	}

	return v
}
