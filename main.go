package main

import (
	"log"
	"proxy/proxy-server"
)

func main() {
	service := proxy_server.NewService()
	endpointFactory := proxy_server.NewEndpointFactory(service)
	_, err := proxy_server.NewServer(endpointFactory)
	if err != nil {
		log.Fatal("SERVER NOT WORK!")
	}
}
