package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type SSH struct {
	Addr   string `mapstructure:"addr"`
	Port   string `mapstructure:"port"`
	Secret string `mapstructure:"secret"`
}

type Mysql struct {
	Host string
	Port string
	Usr  string
	Pwd  string
	Db   string
}

type Oss struct {
	Bucket    string
	SecretId  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
}

// Application 主结构体
type Application struct {
	Name        string
	Port        string
	SshConfig   SSH   `mapstructure:"ssh"`
	MysqlConfig Mysql `mapstructure:"mysql"`
	OssConfig   Oss   `mapstructure:"oss"`
}

var Config Application // 供全局使用

func InitViper() {
	//viper.SetConfigFile("../services/service_user/application.yml")
	viper.SetConfigFile("E:/Go代码页/douyin-project/services/service_user/application.yml")
	err := viper.ReadInConfig() // 读取配置文件
	if err != nil {
		panic(any(fmt.Errorf("viper读取配置文件错误：%v", err)))
	}
	// 监听配置文件，如果被修改会自动重新读取
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err2 := viper.Unmarshal(&Config)
		if err2 != nil {
			panic(any(fmt.Errorf("viper读取配置文件错误：%v", err2)))
		}
	})
	viper.WatchConfig()

	// 将配置映射到结构体
	err1 := viper.Unmarshal(&Config)
	if err1 != nil {
		panic(any(fmt.Errorf("viper读取配置文件错误：%v", err1)))
	}
}
