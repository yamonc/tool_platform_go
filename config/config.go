package config

import "github.com/ltyyz/goprofile"

// 初始化配置文件
func Init() {
	goprofile.Load()
}

// 获取配置文件值
func GetConfig(name string) string {
	return goprofile.GetEnv(name)
}
