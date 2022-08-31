package proxy_server

import "sync"

type ConcurrencyMap interface {
	Insert(key *ProxyRequest, value ProxyResponse)

	List() map[*ProxyRequest]ProxyResponse
}

type concurrencyMap struct {
	mu     sync.RWMutex
	values map[*ProxyRequest]ProxyResponse
}

func NewConcurrencyMap() ConcurrencyMap {
	return &concurrencyMap{
		values: make(map[*ProxyRequest]ProxyResponse),
	}
}

func (c *concurrencyMap) Insert(key *ProxyRequest, value ProxyResponse) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

func (c *concurrencyMap) List() map[*ProxyRequest]ProxyResponse {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result := map[*ProxyRequest]ProxyResponse{}
	for key, value := range c.values {
		result[key] = value
	}
	return result
}
