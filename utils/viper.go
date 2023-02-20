package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type SSH struct {
	Addr   string `mapstructure:"addr"`
	Port   string `mapstructure:"port"`
	Usr    string `mapstructure:"usr"`
	Secret string `mapstructure:"secret"`
}

type Mysql struct {
	Host string
	Port string
	Usr  string
	Pwd  string
	Db   string
}

type Redis struct {
	Addr string
	Port string
	Db   string
}

type Oss struct {
	Bucket    string `mapstructure:"bucket"`
	SecretId  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
}

// Application 主结构体
type Application struct {
	Name        string
	Port        string
	SshConfig   SSH   `mapstructure:"ssh"`
	MysqlConfig Mysql `mapstructure:"mysql"`
	RedisConfig Redis `mapstructure:"redis"`
	OssConfig   Oss   `mapstructure:"oss"`
}

var Config Application // 供全局使用

func InitViper(path string) {
	viper.SetConfigFile(path)
	//viper.AddConfigPath("./")          //设置读取的文件路径
	//viper.SetConfigName("application") //设置读取的文件名
	//viper.SetConfigType("yml")         //设置文件的类型

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
