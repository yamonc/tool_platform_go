package common

import (
	"biligo/constant"
	"fmt"
	"time"
)

// Base 基础结构体，代替 gorm.Model
//
// 使用方式:
//
// type User struct {
// 		common.Model
// }
type Base struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt JSONTime  `json:"createAt"`
	UpdatedAt JSONTime  `json:"updatedAt"`
	DeletedAt *JSONTime `sql:"index" json:"deletedAt"`
}

// JSONTime - 代替 time.Time
type JSONTime time.Time

// MarshalJSON - JSONTime 实现 JSON 序列化方法，格式化成正常的（yyyy-MM-dd HH:mm:ss）日期格式
func (jsonTime JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format(constant.TIME_FORMAT))
	return []byte(stamp), nil
}
