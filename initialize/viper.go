package initialize

import (
	"deliverwater/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var v = viper.New()

func InitConfig() {

	v.SetConfigName("config") //找寻文件的名字
	v.SetConfigType("yaml")   // 找寻文件的类型
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("配置文件未找到！%v\n", err)
			return
		} else {
			panic(fmt.Sprintf("找到配置文件,但是解析错误,%v\n", err))
		}
	}
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		ConfigStruct()
	})
	ConfigStruct()
	log.Print("配置viper服务成功")
}

func ConfigStruct() {
	// 映射到结构体
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("配置重载失败:%s\n", err))
	}

}
