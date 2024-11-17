package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	Server      struct {
		ListenAddr string `mapstructure:"LISTEN_ADDR"`
		Port       string `mapstructure:"PORT"`
	} `mapstructure:"server"`
	Kafka struct {
		Producer struct {
			Brokers []string `mapstructure:"BROKERS"`
			Topics  []string `mapstructure:"TOPICS"`
		} `mapstructure:"producer"`
	} `mapstructure:"kafka"`
}

var v = viper.New()

func New(env string) *AppConfig {
	v.SetConfigName("development")
	v.AddConfigPath("./config")
	v.SetConfigType("toml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	var c AppConfig
	err := v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	return &c
}
