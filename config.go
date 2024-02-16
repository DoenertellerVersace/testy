package main

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	version string
	host    string
	port    int
}

var (
	once    sync.Once
	config  *Config
	version = "0.0.1"
	host    = "localhost"
	port    = 8080
)

func GetConfig() *Config {
	once.Do(func() {
		if p := os.Getenv("PORT"); p != "" {
			port, _ = strconv.Atoi(p)
		}
		if h := os.Getenv("HOST"); h != "" {
			host = h
		}
		if v := os.Getenv("VERSION"); v != "" {
			version = v
		}
		config = &Config{
			version: version,
			host:    host,
			port:    port,
		}
	})
	return config
}
