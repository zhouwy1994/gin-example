package logger

import (
	"fmt"
	"os"
	"time"
	"log"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/zhouwy1994/gin-example/pkg/setting"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var logger *logrus.Logger

func init() {
	var (
		saveDays int
		savePath, baseName string
	)

	sec, err := setting.Cfg.GetSection("logger")
	if err != nil {
		log.Fatal(2, "Fail to get section 'logger': %v", err)
	}

	saveDays = sec.Key("SAVE_DAYS").MustInt(1)
	savePath = sec.Key("SAVE_PATH").MustString("./logs")
	baseName = sec.Key("BASE_NAME").MustString("log")

	logger, err = newLogger(savePath, baseName, saveDays)
	if err != nil {
		log.Fatal(2, "Fail to init logger: %v", err)
	}
}

func newLogger(path, baseFile string, saveDays int) (*logrus.Logger, error) {
	// 如果日志目录不存在则创建
	logger := logrus.New()
	if err := os.MkdirAll(path, 0644); err != nil && !os.IsExist(err) {
		return logger, err
	}
	// 创建日志文件回滚器
	rotate, err := rotatelogs.New(
		// 日志文件名
		path+baseFile+".%Y%m%d%H"+".log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(path+baseFile+".log"),
		// WithRotationTime设置日志分割的时间，设置为24小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),
		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数
		// rotatelogs.WithRotationCount(logSaveDays),
		rotatelogs.WithMaxAge(time.Duration(saveDays)*time.Hour*24),
	)
	if err != nil {
		return logger, err
	}
	// 日志记录等级为InfoLevel以上(包括InfoLevel)
	logger.SetLevel(logrus.InfoLevel)
	// 创建日志Hook，Hook可将日志输出到MQ,ES...此处设置为输出到本地文件,日志内容格式为JSON
	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: rotate,
		logrus.InfoLevel:  rotate,
		logrus.WarnLevel:  rotate,
		logrus.ErrorLevel: rotate,
		logrus.FatalLevel: rotate,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	// 给日志记录器设置Hook
	logger.AddHook(hook)
	// Debug模式,日志输出到屏幕，日志内容格式为TEXT
	logger.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true, ForceColors: true}

	return logger, nil
}

func Debug(class, format string, v ...interface{}) {
	logger.WithFields(logrus.Fields{
		"class": class,
	}).Debug(fmt.Sprintf(format, v...))
}

func Info(class, format string, v ...interface{}) {
	logger.WithFields(logrus.Fields{
		"class": class,
	}).Info(fmt.Sprintf(format, v...))
}

func Warn(class, format string, v ...interface{}) {
	logger.WithFields(logrus.Fields{
		"class": class,
	}).Warn(fmt.Sprintf(format, v...))
}

func Error(class, format string, v ...interface{}) {
	logger.WithFields(logrus.Fields{
		"class": class,
	}).Error(fmt.Sprintf(format, v...))
}

func Fatal(class, format string, v ...interface{}) {
	logger.WithFields(logrus.Fields{
		"class": class,
	}).Fatal(fmt.Sprintf(format, v...))
}
