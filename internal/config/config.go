package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string        `json:"env" yaml:"env" env-required:"true"`
	StoragePath string        `json:"storage_path" yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `json:"token_ttl" yaml:"token_ttl" env-required:"true"`
	GRPC        GRPC          `json:"grpc" yaml:"grpc" env-required:"true"`
}

type GRPC struct {
	Port    int           `json:"port" yaml:"port" env-required:"true"`
	Timeout time.Duration `json:"timeout" yaml:"timeout" env-required:"true"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is required")
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
