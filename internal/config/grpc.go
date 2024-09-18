package config

import (
	"errors"
	"net"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCConfig interface {
	Address() string
}

type Config struct {
	Env        string        `yaml:"env" env-default:"local"`
	GRPC       grpcConfig    `yaml:"grpc"`
	AccessTTL  time.Duration `yaml:"access_ttl" env-default:"1h"`
	RefreshTTL time.Duration `yaml:"refresh_ttl" env-default:"24h"`
}

type grpcConfig struct {
	Port    string        `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

func NewGRPCConfig() (*Config, error) {
	configPath := fetchConfigPath()
	if configPath == "" {
		return nil, errors.New("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("config file does not exist: " + configPath)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, errors.New("config path is empty: " + err.Error())
	}

	return cfg, nil
}

func (cfg *Config) Address() string {
	return net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port)
}

func fetchConfigPath() string {
	var res string

	//flag.StringVar(&res, "config", "", "path to config file")
	//flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
