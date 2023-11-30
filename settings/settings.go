// @Author songxiandong
// @Date 2023/11/29 10:13:00
// @Desc
package settings

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var cfgFile string

func Init() (err error) {
	flag.StringVar(&cfgFile, "c", "", "config file")
	flag.Parse()
	if cfgFile != "" { // 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("sslconfig")
		viper.SetConfigType("yaml")
	}

	err = viper.ReadInConfig()
	if err != nil {
		//读取配置信息失败
		fmt.Println("viperReadInConfig error")
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")

	})

	return nil
}
