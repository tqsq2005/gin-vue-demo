package logging

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/tqsq2005/gin-vue/pkg/file"
	"github.com/tqsq2005/gin-vue/pkg/lfshook"
	"time"
)

type Level int

var (
	AppLog        *log.Logger
	SQLLog        *log.Logger
)

func InitLog() {
	var err error

	//实例化
	AppLog, err = newLog("app")
	if err != nil {
		log.Fatalf("newLog(%s), err:%v\n", "app", err)
	}

	SQLLog, err = newLog("sql")
	if err != nil {
		log.Fatalf("newLog(%s), err:%v\n", "app", err)
	}
}

//生成日志对象
func newLog(logName string) (*log.Logger, error) {
	var err error
	var path string
	filepath := getLogFilePath(logName)
	filename := getLogFileName()
	//创建文件夹
	//生成日志文件
	path, err = file.MustOpen(filename, filepath)
	if err != nil {
		return nil, fmt.Errorf("newLog(%s) err:%v\n", logName, err)
	}
	//path = "/var/log/gin.log"
	// win10共享文件夹会报错：symlink: operation not supported
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),            // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*time.Hour),      // 文件最大保存时间 default: 7 days
		rotatelogs.WithRotationTime(1*time.Hour), // 日志切割时间间隔 默认24小时
	)

	if err != nil {
		return nil, fmt.Errorf("rotatelogs.New err: %v\n", err)
	}

	pathMap := lfshook.WriterMap{
		log.TraceLevel: writer,
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}

	logObject := log.New()
	logObject.Hooks.Add(lfshook.NewHook(
		pathMap,
		&log.JSONFormatter{},
	))
	return logObject, nil
}
