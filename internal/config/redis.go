package config

import (
	"net"
	"os"
	"strconv"
)

type RedisConfig struct {
	host     string
	port     string
	Password string
	DBNum    int
}

func NewRedisConfig() *RedisConfig {
	dbnum, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		dbnum = 0
	}
	return &RedisConfig{
		host:     os.Getenv("REDIS_HOST"),
		port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DBNum:    dbnum,
	}
}

func (cfg RedisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
