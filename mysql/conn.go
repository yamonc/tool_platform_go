package mysql

import (
	"biligo/config"
	"biligo/constant"
	"biligo/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

// 数据库类型
const DatabaseType = "mysql"

// 获取mysql连接字符串
func getMySQLConnectionUrl() string {
	var url = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf(url,
		config.GetConfig(constant.DbUsername),
		config.GetConfig(constant.DbPassword),
		config.GetConfig(constant.DbHost),
		config.GetConfig(constant.DbDatabase))
}

// 初始化数据库连接
func Init() {
	log.Debug("初始化数据库连接..")
	url := getMySQLConnectionUrl()
	db, err := gorm.Open(DatabaseType, url)
	if err != nil {
		log.Panic("初始化数据库连接失败", err)
	}
	db.SingularTable(true)

	Conn = db
}

// 关闭数据库连接池
func Close() {
	err := Conn.Close()
	if err != nil {
		log.Error("关闭数据库连接失败", err)
	}
}

// 检查 sql 结果
func CheckResult(result *gorm.DB) error {

	return nil
}
