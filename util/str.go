package util

import (
	"biligo/log"
	"regexp"
	"strconv"

	"github.com/gofrs/uuid"
)

const (
	REGEXP_INT = "^[0-9]+$"
)

// UUID - 生成一个UUID
func UUID() string {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Error("failed to generate UUID: %v", err)
	}
	return u4.String()
}

// StrMatch - 字符串匹配正则
func StrMatch(reg string, s string) bool {
	ok, _ := regexp.Match(reg, []byte(s))
	return ok
}

// IsInt - 判断一个字符串是否为整数
func IsInt(s string) bool {
	return StrMatch(REGEXP_INT, s)
}

// ToInt - 将字符串转成 int，忽略错误
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Error("转换成 int 失败: %v", err)
	}
	return i
}
