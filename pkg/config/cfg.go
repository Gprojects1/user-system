package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Port int
	}
	DB struct {
		Host     string
		Username string
		Password string
		Database string
	}
}

var Defaults = map[string]interface{}{
	"http": map[string]string{
		"port": "8080",
	},
	"db": map[string]string{
		"host":     "dbPgsql:5432",
		"username": "root",
		"password": "love",
		"database": " user-service-db",
	},
}

func Read(appName string, defaults map[string]interface{}, cfg interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(appName)
	v.AddConfigPath("/etc/")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if cfg != nil {
		err := v.Unmarshal(cfg)
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}
