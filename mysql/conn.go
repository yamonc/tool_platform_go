package mysql

import (
	"biligo/config"
	"biligo/constant"
	"biligo/log"
	"biligo/modules/app/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic("初始化数据库连接失败", err)
	}
	// 开启数据库自映射，直接通过结构体映射成为数据库表对象。注：不支持结构体删除字段同步表结构
	errMigrate := db.AutoMigrate(&model.Car{})
	logger.Default.LogMode(logger.Info)
	if errMigrate != nil {
		return
	}
	Conn = db
}

// 关闭数据库连接池
func Close() {
	// 最新版本 gorm v2.0 已经没有 Close() 函数了 Sad :(

	// err := Conn.Close()
	// if err != nil {
	// 	log.Error("关闭数据库连接失败", err)
	// }
}

// 检查 sql 结果
func CheckResult(result *gorm.DB) error {

	return nil
}
