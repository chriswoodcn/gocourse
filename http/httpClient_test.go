package http

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
)

func sampleRequest() {
	resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		fmt.Println("http get error")
	}
	defer func(resp *http.Response) {
		_ = resp.Body.Close()
	}(resp)
	response, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s ", response)
}
func customRequest() {
	request, _ := http.NewRequest(http.MethodGet,
		"https://www.imooc.com",
		nil,
	)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	response, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s ", response)
}
func TestHttpCli(t *testing.T) {
	customRequest()
}
