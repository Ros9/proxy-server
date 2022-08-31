package proxy_server

import (
	uuid "github.com/nu7hatch/gouuid"
	"net/http"
)

type Service interface {
	DoProxy(proxyRequest ProxyRequest) (*ProxyResponse, error)
	GetList() []RequestResponse
}

type service struct {
	httpClt *http.Client
	reqResp ConcurrencyMap
}

func NewService() Service {
	httpClt := http.DefaultClient
	return &service{
		httpClt: httpClt,
		reqResp: NewConcurrencyMap(),
	}
}

func (s *service) DoProxy(proxyRequest ProxyRequest) (*ProxyResponse, error) {
	httpClt := http.DefaultClient
	req, _ := http.NewRequest(proxyRequest.Method, proxyRequest.Url, nil)
	for key, val := range proxyRequest.Headers {
		req.Header.Add(key, val)
	}
	resp, err := httpClt.Do(req)
	if err != nil {
		return nil, err
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	proxyResponse := ProxyResponse{
		Id:      uid.String(),
		Status:  resp.Status,
		Headers: make(map[string][]string),
		Length:  int(resp.ContentLength),
	}
	for key, val := range resp.Header {
		proxyResponse.Headers[key] = val
	}
	s.reqResp.Insert(&proxyRequest, proxyResponse)
	return &proxyResponse, err
}

func (s *service) GetList() []RequestResponse {
	requestsResponses := []RequestResponse{}
	reqRespList := s.reqResp.List()
	for key, val := range reqRespList {
		requestsResponses = append(requestsResponses, RequestResponse{
			ProxyRequest:  *key,
			ProxyResponse: val,
		})
	}
	return requestsResponses
}
