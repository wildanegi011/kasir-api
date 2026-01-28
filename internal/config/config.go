package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
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

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetConfigFile(".env")
	_ = v.ReadInConfig()

	v.SetDefault("app.name", v.GetString("APP_NAME"))
	v.SetDefault("app.port", v.GetString("APP_PORT"))
	v.SetDefault("database.url", v.GetString("DATABASE_URL"))
	v.SetDefault("database.max_open_conns", v.GetInt("DATABASE_MAX_OPEN_CONNS"))
	v.SetDefault("database.max_idle_conns", v.GetInt("DATABASE_MAX_IDLE_CONNS"))

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Println("Failed to unmarshal config:", err)
	}

	return &config
}

func GetConfig() *Config {
	once.Do(func() {
		cfg = LoadConfig()
	})
	return cfg
}
