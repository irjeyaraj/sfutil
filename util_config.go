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
	"fmt"
	"runtime"

	"github.com/BurntSushi/toml"
)

var Config Sedconfig

const ConfFileName = "sfutil.conf"

type server struct {
	Port       int
	Protocol   string
	Hostname   string
	ServerCert string
	ServerKey  string
}

type sedsys struct {
	LogLocation string
	LogLevel    int32
}

type Sedconfig struct {
	Server server
	Log    sedsys
}

// Parse the config file and return data using Sedconfig data structure
func ParseSystemConfig(conffile string, appName string) Sedconfig {
	var confloc string
	var Config Sedconfig

	if runtime.GOOS == "windows" {
		confloc = "etc" + Path_separator() + conffile
	} else {
		conpath01 := "etc" + Path_separator() + appName + Path_separator() + conffile
		conpath02 := Path_separator() + "home" + Path_separator() + CurrentUser() +
			Path_separator() + "." + appName + "/" + conffile
		if File_is_exists(conpath01) {
			confloc = conpath01
		} else if File_is_exists(conpath02) {
			confloc = conpath02
		} else {
			confloc = Path_separator() + conpath01
		}
	}

	_, err := toml.DecodeFile(confloc, &Config)
	if err != nil {
		errText := fmt.Sprint(err)
		fmt.Println(errText)
	}

	return Config
}
