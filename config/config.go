package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	CONFIG_FILE_ENV_NAME = "CONFIGPATH"
)

var ENV_PATH_TYPES_AVAILABLES = map[string]string{
	"local":  "./config/config-local",
	"docker": "./config/config-docker",
}

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	AppVersion string
	Port       string
	Mode       string
	Debug      bool
}

type LoggerConfig struct {
	Development        bool
	DisableCaller      bool
	DisableStackTRACER bool
	Level              string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config *Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetConfigPath(env string) (string, error) {
	if value, exists := os.LookupEnv(CONFIG_FILE_ENV_NAME); exists {
		return value, nil
	} else {
		if value, ok := ENV_PATH_TYPES_AVAILABLES[env]; ok {
			return value, nil
		} else {
			return nil, errors.New(fmt.Sprintf())
		}
	}

}
