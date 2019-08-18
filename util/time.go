package util

import (
	"biligo/log"
	"time"
)

// 获取系统当前毫秒时间戳
func CurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 计算叠加后的日期
func AddTime(target time.Time, duration string) time.Time {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		log.Error(err)
	}
	return target.Add(timeDuration)
}
