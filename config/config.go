package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Database *Database `mapstructure:"database" validate:"required"`
}

type Database struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DBName   string `mapstructure:"dbname" validate:"required"`
	SSLMode  string `mapstructure:"sslmode" validate:"required"`
	Schema   string `mapstructure:"schema" validate:"required"`
}

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
		validate := validator.New()
		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})
	fmt.Println("configInstance was loaded successfully")
	return configInstance
}
