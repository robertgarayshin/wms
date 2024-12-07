package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	Log  `yaml:"logger"`
	PG   `yaml:"postgres"`
}

type App struct {
	Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
}

type PG struct {
	PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	URL     string `env-required:"true"                 env:"PG_URL"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig(path+"/config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
