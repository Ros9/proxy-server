package proxy_server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
}

type server struct {
	router          *gin.Engine
	endpointFactory EndpointFactory
}

func NewServer(endpointFactory EndpointFactory) (Server, error) {
	router := gin.Default()
	router.Handle("POST", "/proxy-server", endpointFactory.DoProxyEndpoint())
	router.Handle("GET", "/proxy-server", endpointFactory.GetListEndpoint())
	err := router.Run(":8080")
	if err != nil {
		return nil, err
	}
	return &server{
		router:          gin.Default(),
		endpointFactory: endpointFactory,
	}, nil
}
