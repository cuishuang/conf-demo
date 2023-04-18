package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/configor"
)

type Config struct {
	APPName string `default:"app name"`
	DB      struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}

func main() {
	var conf = Config{}

	// reload模式,可实现热加载

	err := configor.New(&configor.Config{
		AutoReload:         true,
		AutoReloadInterval: time.Second,
		AutoReloadCallback: func(config interface{}) {
			// config发生变化后出发什么操作
			fmt.Printf("配置文件发生了变更%#v\n", config)
		},
	}).Load(&conf, "config.yml")

	// 无reload模式
	//err := configor.Load(&conf, "config.yml")

	// err := configor.New(&configor.Config{Debug: true}).Load(&conf, "config.yml")  // 测试模式，也可以通过环境变量开启测试模式(CONFIGOR_DEBUG_MODE=true go run main.go )，这样就无需修改代码

	//err := configor.New(&configor.Config{Verbose: true}).Load(&conf, "config.yml") // 模式，也可以通过环境变量开启详细模式(CONFIGOR_VERBOSE_MODE=true go run main.go )，这样就无需修改代码
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", conf)

	time.Sleep(100000e9)
}
