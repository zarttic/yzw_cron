package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"yzw_cron/logger"
)

func File(body []byte) error {
	dirPath := "craw/" + TimeStr("2006-01-02")
	filePath := TimeStr("2006-01-02-15-04")
	fmt.Println(dirPath)
	filePath = filepath.Join(dirPath, filePath+".txt")
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		logger.LOGGER.Println(fmt.Sprintf("创建文件夹{%s}成功", dirPath))
		if err != nil {
			logger.LOGGER.Println(fmt.Sprintf("创建文件夹{%s}失败", dirPath))
			return err
		}
	}
	err = os.WriteFile(filePath, []byte(body), 0644)
	if err != nil {
		logger.LOGGER.Println(fmt.Sprintf("写入文件{%s}失败", filePath))
		return err
	}
	//io.Writer()
	logger.LOGGER.Println(fmt.Sprintf("写入文件{%s}成功", filePath))
	return nil
}
