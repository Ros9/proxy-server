package proxy_server

import (
	"github.com/gin-gonic/gin"
)

type EndpointFactory interface {
	DoProxyEndpoint() gin.HandlerFunc
	GetListEndpoint() gin.HandlerFunc
}

type endpointFactory struct {
	service Service
}

func NewEndpointFactory(service Service) EndpointFactory {
	return &endpointFactory{service: service}
}

func (ef *endpointFactory) DoProxyEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyRequest := ProxyRequest{}
		if err := c.ShouldBindJSON(&proxyRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		proxyResp, err := ef.service.DoProxy(proxyRequest)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, proxyResp)
	}
}

func (ef *endpointFactory) GetListEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestsResponses := ef.service.GetList()
		c.JSON(200, requestsResponses)
	}
}
