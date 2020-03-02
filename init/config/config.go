package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Gin `json:"gin"`
	Mysql `json:"mysql"`
	Redis `json:"redis"`
}

type Gin struct {
	Mode string `json:"mode"`
}

type Mysql struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP string `json:"ip"`
	Port string `json:"port"`
	DBName string `json:"dbname"`
	Param string `json:"param"`
}
type Redis struct {
	IP string `json:"ip"`
	Port string `json:"port"`
	DB int `json:"db"`
}

var GlobalConfig Config
var GlobalViper *viper.Viper

func Init() {
	v:=viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./static/config/")
	v.SetConfigType("json")
	err:=v.ReadInConfig()
	if err!= nil{
		panic(fmt.Sprintf("读取配置文件出错：%v",err.Error()))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置已改变：",e.Name)
		err=v.Unmarshal(&GlobalConfig)
		if err!=nil{
			fmt.Printf("解析配置文件出错：%v\n",err.Error())
		}
	})

	err=v.Unmarshal(&GlobalConfig)
	if err!=nil{
		panic(fmt.Sprintf("解析配置文件出错：%v",err.Error()))
	}
	GlobalViper=v
	fmt.Println("配置文件已加载完成...")
}
