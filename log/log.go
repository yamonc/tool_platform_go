package log

import "github.com/sirupsen/logrus"

// 就算使用 logrus api 也要包一层

type Level logrus.Level

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// 设置日志级别
func SetLevel(level Level) {
	switch level {
	case TraceLevel:
		logrus.SetLevel(logrus.TraceLevel)
	case DebugLevel:
		logrus.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		logrus.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		logrus.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		logrus.SetLevel(logrus.ErrorLevel)
	case FatalLevel:
		logrus.SetLevel(logrus.FatalLevel)
	case PanicLevel:
		logrus.SetLevel(logrus.PanicLevel)
	}
}

func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Println(args ...interface{}) {
	logrus.Println(args...)
}

func Print(args ...interface{}) {
	logrus.Print(args...)
}
