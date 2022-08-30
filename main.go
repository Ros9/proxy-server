package main

import (
	"log"
	"proxy/proxy"
)

func main() {
	service := proxy.NewService()
	endpointFactory := proxy.NewEndpointFactory(service)
	_, err := proxy.NewServer(endpointFactory)
	if err != nil {
		log.Fatal("SERVER NOT WORK!")
	}
}
