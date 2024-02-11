package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"yzw_cron/config"
)

var LOGGER *log.Logger

func init() {
	logDirPath := "log"
	//指定路径的文件，无则创建
	_, err := os.Stat(logDirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(logDirPath, 0755)
		if err != nil {
			fmt.Println("创建日志目录失败")
		}
	}
	logFile, err := os.OpenFile(filepath.Join(logDirPath, config.Conf.LogConfig.Filename), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	LOGGER = log.New(logFile, config.Conf.LogConfig.Prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
