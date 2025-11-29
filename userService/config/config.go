package config

import "sync"

var (
	instance *Config
	once     sync.Once
	initErr  error
)

type Config struct {
	Postgresql  Postgresql
	Server      Server
	Application Application
}

func GetInstance() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		if initErr != nil {
			instance = nil
		}
	})
	return instance, initErr
}
