package config

import (
	"sync"
)

type Config struct {
	MongoDB struct {
		Host       string `yaml:"host" env-required:"true"`
		Port       string `yaml:"port" env-required:"true"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		AuthDB     string `yaml:"auth_db" env-required:"true"`
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
		//if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
		//help, _ := cleanenv.GetDescription(instance, nil)
		//logger.Info(help)
		//logger.Fatal(err)
		//}
	})
	return instance

}
