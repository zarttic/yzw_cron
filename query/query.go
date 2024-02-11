package query

import (
	"io"
	"net/http"
	"net/url"
	"yzw_cron/config"
	"yzw_cron/logger"
	"yzw_cron/utils"
)

func Query() {
	params := url.Values{}
	params.Set("xm", config.Conf.YzwConfig.Name)
	params.Set("zjhm", config.Conf.YzwConfig.ID)
	params.Set("ksbh", config.Conf.YzwConfig.StudentCode)
	params.Set("bkdwdm", config.Conf.YzwConfig.SchoolCode)

	Url, err := url.Parse(config.Conf.YzwConfig.Url)
	if err != nil {
		logger.LOGGER.Println("url错误:", err)
		return
	}
	Url.RawQuery = params.Encode()
	// 全局默认客户端
	client := http.DefaultClient
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		logger.LOGGER.Println("创建请求失败:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.LOGGER.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体到一个切片中
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LOGGER.Println("读取响应失败:", err)
		return
	}

	//fmt.Println(string(body))

	err = utils.File(body)
	if err != nil {
		logger.LOGGER.Println("写入文件失败:", err)
		return
	}
	logger.LOGGER.Println("写入成功")
	match, err := utils.RegMatch("无查询结果", string(body))
	if err != nil {
		logger.LOGGER.Println("正则匹配失败:", err)
	}
	if !match {
		err = utils.SendEmail(string(body))
		if err != nil {
			logger.LOGGER.Println("发送邮件失败:", err)
			return
		}
		logger.LOGGER.Println("发送邮件成功:")
	}
}
