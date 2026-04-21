package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env          string   `mapstructure:"env"`
	KafkaBrokers []string `mapstructure:"kafka_brokers"`
	Port         int      `mapstructure:"port"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetEnvPrefix("TOPK")
	viper.AutomaticEnv()

	cfg := &Config{
		Env:          viper.GetString("env"),
		KafkaBrokers: viper.GetStringSlice("kafka_brokers"),
		Port:         viper.GetInt("port"),
	}

	if cfg.Env == "dev" {
		cfg.KafkaBrokers = []string{"kafka:9092"}
		cfg.Port = 8080
	}
	log.Printf("Loaded config for env=%s", cfg.Env)
	return cfg
}
