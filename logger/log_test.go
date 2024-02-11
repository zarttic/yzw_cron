package logger

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"yzw_cron/utils"
)

func TestInit(t *testing.T) {
	logDirPath := "log"
	logFilePath := utils.TimeStr("2006-01-02")
	//指定路径的文件，无则创建
	_, err := os.Stat(logDirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(logDirPath, 0755)
		if err != nil {
			t.Errorf("创建日志失败: %v", err)
		}
	}
	logFile, err := os.OpenFile(filepath.Join(logDirPath, logFilePath+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Errorf("打开日志文件失败: %v", err)
	}
	LOGGER = log.New(logFile, "[yzw_info]", log.Ldate|log.Ltime|log.Lshortfile)
	if LOGGER == nil {
		t.Errorf("LOGGER初始化失败")
	}
}
