package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	Mongo  Mongo
}

type Server struct {
	Host string
	Port string
}

type Mongo struct {
	Host string
	Port string
}

func ReadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error whilw reading config: %v\n", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error while unmarshalling: %v\n", err)
	}

	return &cfg
}
