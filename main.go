package main

import (
	"biligo/config"
	"biligo/constant"
	"biligo/log"
	"biligo/modules"
	"biligo/mysql"
)

// modules 程序入口
func main() {
	// 配置注册
	config.Init()

	// 设置日志级别
	log.SetLevel(log.InfoLevel)

	// 连接数据库
	mysql.Init()
	defer mysql.Close()

	// 注册路由并启动
	err := modules.RegisterRouter(modules.NewRouter()).
		Run(":" + config.GetConfig(constant.HttpServerPort))

	if err != nil {
		log.Panic(err)
	}
}
