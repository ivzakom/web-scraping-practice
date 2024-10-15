package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Listen struct {
		Type   string `yaml:"type" env-default:"tcp"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"1234"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `yaml:"host" env-required:"true" env-default:"localhost"`
		Port       string `yaml:"port" env-required:"true" env-default:"27017"`
		Username   string `yaml:"username" `
		Password   string `yaml:"password" `
		AuthDB     string `yaml:"auth_db" `
		Database   string `yaml:"database" env-required:"true"`
		Collection string `yaml:"collection" env-required:"true"`
	} `yaml:"mongodb" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {

	once.Do(func() {
		//logger := logging.GetLogger()
		//logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(help)
			//logger.Info(help)
			//logger.Fatal(err)
		}
	})
	return instance

}
