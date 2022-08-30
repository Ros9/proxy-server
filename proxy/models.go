package proxy

type ProxyRequest struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	Id      string              `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}

type RequestResponse struct {
	ProxyRequest  ProxyRequest  `json:"proxy_request"`
	ProxyResponse ProxyResponse `json:"proxy_response"`
}
