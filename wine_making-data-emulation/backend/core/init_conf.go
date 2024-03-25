package core

import (
	"backend/config"
	"backend/gobal"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const ConfigYaml = "settings.yaml"

func InitConfig() *config.Config {
	data, err := os.ReadFile(ConfigYaml)
	if err != nil {
		log.Fatal("read config file error")
		return nil
	}

	config := &config.Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatal("init config error")
	}
	//fmt.Println(config)
	return config
}

func SetConfig() error {
	data, err := yaml.Marshal(gobal.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigYaml, data, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
