package config

import "time"

type Application struct {
	Version     string        `env:"VERSION"`
	Environment string        `env:"ENVIRONMENT"`
	JWTSecret   string        `env:"JWT_SECRET"`
	JWTExp      time.Duration `env:"JWT_EXPIRES"`
}
