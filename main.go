package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	serverChan := make(chan ServerEvent, 10)

	config := GetConfig()
	versionServer := http.NewServeMux()
	versionServer.HandleFunc("/", getVersionHandler(config, serverChan))
	go serveVersion(config, versionServer, serverChan)

	helloServer := http.NewServeMux()
	helloServer.HandleFunc("/", getHelloHandler(serverChan))
	go serveHelloWorld(config, helloServer, serverChan)

	for e := range serverChan {
		if e.err != nil {
			log.Println(e.msg, e.err)
		} else {
			log.Println(e.msg)
		}
	}
}

func getHelloHandler(serverChan chan ServerEvent) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i, err := w.Write([]byte("hello, world"))
		if err != nil {
			serverChan <- ServerEvent{msg: "failed to write response", err: err}
		} else {
			serverChan <- ServerEvent{msg: "wrote " + strconv.Itoa(i) + " bytes"}
		}
	}
}

func getVersionHandler(config *Config, c chan<- ServerEvent) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i, err := w.Write([]byte("version: " + config.version))
		if err != nil {
			c <- ServerEvent{msg: "failed to write response", err: err}
		} else {
			c <- ServerEvent{msg: "wrote " + strconv.Itoa(i) + " bytes"}
		}
	}
}

func serveVersion(config *Config, mux *http.ServeMux, c chan<- ServerEvent) {
	defer close(c)
	c <- ServerEvent{msg: "server starting"}
	err := http.ListenAndServe(config.host+":"+strconv.Itoa(config.port), mux)
	if err != nil {
		c <- ServerEvent{msg: "server went down", err: err}
	}
}

func serveHelloWorld(config *Config, mux *http.ServeMux, c chan<- ServerEvent) {
	defer close(c)
	c <- ServerEvent{msg: "server starting"}
	err := http.ListenAndServe(config.host+":"+strconv.Itoa(config.port+1), mux)
	if err != nil {
		c <- ServerEvent{msg: "server went down", err: err}
	}
}
