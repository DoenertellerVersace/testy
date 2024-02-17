package main

import (
	"log"
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
		log.Printf("port: %d\n", port)

		if h := os.Getenv("HOST"); h != "" {
			host = h
		}
		log.Printf("host: %s\n", host)

		if v := os.Getenv("VERSION"); v != "" {
			version = v
		}
		log.Printf("version: %s\n", version)

		config = &Config{
			version: version,
			host:    host,
			port:    port,
		}
	})
	return config
}
