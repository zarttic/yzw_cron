package main

import (
	"github.com/robfig/cron/v3"
	"yzw_cron/config"
	_ "yzw_cron/config"
	"yzw_cron/logger"
	"yzw_cron/query"
)

// 打包前match别忘改
func main() {
	logger.LOGGER.Println("服务启动")
	c := cron.New()
	c.AddFunc(config.Conf.CronConfig.Spec, query.Query)
	c.Run()
}
