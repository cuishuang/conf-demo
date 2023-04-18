package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	loadConfig()
}

func loadConfig() {

	configVar := "shuang-config.yaml"
	configVar = "" // 这行如果注释掉，则从指定的configVar读取配置文件；否则就各种条件去找了

	if configVar != "" {
		// SetConfigFile 显式定义配置文件的路径、名称和扩展名。
		// Viper 将使用它而不检查任何配置路径。
		viper.SetConfigFile(configVar)
	} else {

		// 如果没有显式指定配置文件，则

		// 会去下面的路径里找文件名`cui-config`的文件  name of config file (without extension)
		// 按照 []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}的顺序(居然还支持Java用的properties)
		viper.SetConfigName("cui-config")
		viper.AddConfigPath("/etc/myapp") // 找寻的路径
		viper.AddConfigPath("$HOME/.myapp/")
		viper.AddConfigPath(".")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件 %s 发生了更改!!! 最新的Global.Source这个字段的值为 %s:", e.Name, viper.GetString("Global.Source"))
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	fmt.Printf("到底用的是哪个配置文件: '%s'\n", viper.ConfigFileUsed())

	fmt.Printf("Global.Source这个字段的值为: '%s'\n", viper.GetString("Global.Source"))

	time.Sleep(10000e9)
}
