package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	JwtKey string

	SavePath string
	FileSize int64
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误")
	}

	LoadServer(file)
	LoadFs(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("123456789")
}

func LoadFs(file *ini.File) {
	SavePath = file.Section("fs").Key("SavePath").MustString("./tmp/")
	FileSize = file.Section("fs").Key("FileSize").MustInt64(10 * 1024 * 1024)
}
