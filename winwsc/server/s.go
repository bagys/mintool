package server

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type myconfig []map[string]string

type ServiceConfig struct {
	Svc myconfig `mapstructure:"services"`
}

func NewServiceConfig() *ServiceConfig {
	return config()
}

func config() *ServiceConfig {

	filep := fmt.Sprintf("%s/config.json", GetRunPath2())
	_, err := os.Stat(filep)
	if err != nil {
		filep = fmt.Sprintf("./config.json")
		_, err = os.Stat(filep)
		if err != nil {
			log.Fatal(err)
		}
	}

	viper.SetConfigFile(filep) // 指定配置文件路径
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("配置文件没有找到，" + err.Error())
			// Config file not found; ignore error if desired
		} else {
			log.Fatal("配置文件解析错误，" + err.Error())
		}
	}

	Services := &ServiceConfig{}
	configErr := viper.Unmarshal(Services)
	if configErr != nil {
		log.Fatal(configErr)
	}

	return Services
}

func (this ServiceConfig) List() myconfig {
	return this.Svc
}
