package util

import (
	"time"
)

// 获取系统当前毫秒时间戳
func CurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
