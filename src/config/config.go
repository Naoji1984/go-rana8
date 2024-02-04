package config

import "github.com/caarlos0/env/v10"

var Config *AppConfig

type AppConfig struct {
	RDB  *RDBConfig
	Port string `env:"PORT" envDefault:"8080"`
}

type RDBConfig struct {
	Host     string `env:"RANA8_DB_HOST"`
	Port     string `env:"RANA8_DB_PORT"`
	User     string `env:"RANA8_DB_USER"`
	Password string `env:"RANA8_DB_PASSWORD"`
	Name     string `env:"RANA8_DB_NAME"`
}

func LoadConfig() {
	Config = &AppConfig{}
	env.Parse(Config)
}
