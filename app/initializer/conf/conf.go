package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	Conf      *appConf
	confMutex sync.Mutex
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic("配置读取失败：" + err.Error())
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic("配置解码失败: " + err.Error())
	}
	// 监听变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if e.Name == viper.ConfigFileUsed() {
			confMutex.Lock()
			defer confMutex.Unlock()
			fmt.Println("配置文件已变更，正在重新加载...")
			tmpConf := &appConf{}
			if err := viper.Unmarshal(tmpConf); err != nil {
				fmt.Println("配置重新加载出错: " + err.Error())
			}
			Conf = tmpConf
		}
	})
}

type appConf struct {
	Server ServerConf `mapstructure:"server"`
	Log    LogConf    `mapstructure:"log"`
	DB     DBConf     `mapstructure:"db"`
	Redis  RedisConf  `mapstructure:"redis"`
}
