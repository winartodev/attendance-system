package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is struct to store all config from yaml file
type Config struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"app"`

	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	SMTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"smtp"`

	AMQP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"amqp"`
}

// NewConfig will read file and return config struct
func (c Config) NewConfig() Config {

	var filepath string
	if os.Getenv("APP_ENV") == "production" {
		filepath = "./files/config.production.yaml"
	} else {
		filepath = "./files/config.development.yaml"
	}

	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("[Config] [NewConfig] fail to read file %v", err)
		return c
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("[Config] [NewConfig] unmarshall error %v", err)
		return c
	}

	return c
}
