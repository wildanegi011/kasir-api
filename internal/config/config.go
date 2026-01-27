package config

import (
	"log"
	"os"
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

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("app.port", "8080")
	v.SetDefault("database.max_open_conns", 10)
	v.SetDefault("database.max_idle_conns", 5)

	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		v.AddConfigPath("./internal/config")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		_ = v.ReadInConfig()
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal("config: failed to unmarshal:", err)
	}

	return &config
}

func GetConfig() *Config {
	once.Do(func() {
		cfg = LoadConfig()
	})
	return cfg
}
