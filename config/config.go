package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(UnitConfig)

type MailConfig struct {
	From     string   `mapstructure:"from"`
	To       []string `mapstructure:"to"`
	Subject  string   `mapstructure:"subject"`
	Addr     string   `mapstructure:"addr"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	Host     string   `mapstructure:"host"`
	Identity string   `mapstructure:"identify"`
}
type YzwConfig struct {
	Name        string `mapstructure:"xm"`
	ID          string `mapstructure:"zjhm"`
	StudentCode string `mapstructure:"ksbh"`
	SchoolCode  string `mapstructure:"bkdwdm"`
	Url         string `mapstructure:"url"`
}
type LogConfig struct {
	Filename string `mapstructure:"filename"`
	Prefix   string `mapstructure:"prefix"`
}
type CronConfig struct {
	Spec string `mapstructure:"spec"`
}

type UnitConfig struct {
	*MailConfig `mapstructure:"mail"`
	*YzwConfig  `mapstructure:"yzw"`
	*LogConfig  `mapstructure:"log"`
	*CronConfig `mapstructure:"cron"`
}

func init() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取配置失败%v", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件被修改")
		err := viper.Unmarshal(&Conf)
		if err != nil {
			panic(fmt.Sprintf("配置文件修改出错%v", err))
		}
	})
	// 读取修改后的配置文件
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取修改配置失败%v", err))
	}
	//将配置文件写入到conf结构体
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Sprintf("配置文件写入结构体失败%v", err))
	}

}
