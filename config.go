/*
 * util_config.go
 *
 * Copyright 2022 Immanuel Jeyaraj
 *
 * Author: Immanuel Jeyaraj <irj@sefier.com>
 *
 * Created date: 3 June 2019
 */

package sfutil

import (
	"github.com/spf13/viper"
)

var ConfigData map[string]interface{}

func ReadConfig(confFile string, confType string, appName string) (map[string]interface{}, error) {
	viper.SetConfigName(confFile)
	viper.SetConfigType(confType)
	viper.AddConfigPath("/etc/" + appName + "/")
	viper.AddConfigPath("$HOME/." + appName + "/")
	viper.AddConfigPath("./etc")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	ConfigData = viper.AllSettings()
	return ConfigData, err
}
