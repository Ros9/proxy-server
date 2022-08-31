package proxy_server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDoProxy(t *testing.T) {
	data := []byte(`{
    	"method": "GET",
    	"url": "https://gobyexample.com/testing",
    	"headers": {
        	"content-type": "text/html"
    	}
	}`)
	req, _ := http.NewRequest("POST", "http://localhost:8080/proxy", bytes.NewReader(data))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("Making 'POST http://localhost:8080/proxy' request failed!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Making 'POST http://localhost:8080/proxy' request failed!")
	}
	proxyResp := ProxyResponse{}
	err = json.Unmarshal(body, &proxyResp)
	if err != nil {
		t.Fatal("Making 'POST http://localhost:8080/proxy' request failed!")
	}
	wantStatus := "200 OK"
	if wantStatus != proxyResp.Status {
		t.Fatalf("statuses is different want = %s and result = %s", wantStatus, proxyResp.Status)
	}
}

func TestDoProxy400(t *testing.T) {
	data := []byte(`{
    	"method": "GET",
    	"url": "https://gobyexample.com/qwerty",
    	"headers": {
        	"content-type": "text/html"
    	}
	}`)
	req, _ := http.NewRequest("POST", "http://localhost:8080/proxy", bytes.NewReader(data))
	resp, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Making 'POST http://localhost:8080/proxy' request failed!")
	}
	proxyResp := ProxyResponse{}
	err = json.Unmarshal(body, &proxyResp)
	if err != nil {
		t.Fatal("Making 'POST http://localhost:8080/proxy' request failed!")
	}
	wantStatus := "404 Not Found"
	if wantStatus != proxyResp.Status {
		t.Fatalf("statuses is different want = %s and result = %s", wantStatus, proxyResp.Status)
	}
}
