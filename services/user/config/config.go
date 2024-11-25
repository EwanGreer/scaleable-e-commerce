package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	Server      struct {
		ListenAddr string `mapstructure:"LISTEN_ADDR"`
		Port       string `mapstructure:"PORT"`
	} `mapstructure:"server"`
	Database struct {
		ConnectionString string `mapstructure:"CONNECTION_STRING"`
	} `mapstructure:"database"`
	Kafka struct {
		Producer struct {
			Brokers []string `mapstructure:"BROKERS"`
			Topics  []string `mapstructure:"TOPICS"`
		} `mapstructure:"producer"`
	} `mapstructure:"kafka"`
}

var v = viper.New()

func Load() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	v.SetConfigName(os.Getenv("ENV"))
	v.AddConfigPath("./config")
	v.SetConfigType("toml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	var c AppConfig
	err = v.Unmarshal(&c)
	if err != nil {
		log.Panic(err)
	}

	return &c
}
