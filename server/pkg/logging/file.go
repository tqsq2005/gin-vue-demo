package logging

import (
	"fmt"
	. "github.com/tqsq2005/gin-vue/pkg/setting"
	"strings"
	"time"
)

// getLogFilePath: get the log file save path
func getLogFilePath(extraPath string) string {
	//先判断是否有"/"，没有则加上去
	if !strings.HasSuffix(extraPath, "/") {
		extraPath = extraPath + "/"
	}
	return fmt.Sprintf("%s%s%s",
		Config.App.RuntimeRootPath,
		Config.Log.LogSavePath,
		extraPath,
	)
}

// getLogFileName: get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		Config.Log.LogSaveName,
		time.Now().Format(Config.Log.TimeFormat),
		Config.Log.LogFileExt,
	)
}
