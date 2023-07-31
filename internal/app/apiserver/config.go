package apiserver

import "github.com/Dor1ma/Basic-http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
