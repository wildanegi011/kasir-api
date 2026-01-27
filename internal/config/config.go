package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	URL          string `mapstructure:"url"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	v := viper.New()

	v.SetEnvPrefix("KASIR")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("app.port", "8080")
	v.SetDefault("database.max_open_conns", 10)
	v.SetDefault("database.max_idle_conns", 5)

	v.AddConfigPath("./internal/config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Println("failed to read config", err)
		return nil
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Println("failed to unmarshal config", err)
		return nil
	}

	return &config
}

func GetConfig() *Config {
	once.Do(func() {
		cfg = LoadConfig()
	})
	return cfg
}
