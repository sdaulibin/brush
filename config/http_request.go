package config

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpRequest struct {
	Url    string
	Method string
	Header *http.Header
}

type Params struct {
	Params map[string]string
}

func NewDefaultHttpRequest(token, url, method string, params *Params) *HttpRequest {
	return &HttpRequest{
		Url:    url + PrepareParams(params.Params),
		Method: method,
		Header: initHeader(token),
	}
}

func initHeader(token string) *http.Header {
	return &http.Header{
		"Host":            {"ecosys-web.china-inv.cn"},
		"Accept-Encoding": {"ecosys-web.china-inv.cn"},
		"sign":            {"fbbee325b38811c5bf83c06df62851e2"},
		"Referer":         {"https://ecosys-web.china-inv.cn/index.html"},
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"},
		"Connection":      {"Keep-Alive"},
		"Content-Type":    {"application/json"},
		"accessToken":     {token},
	}
}

func PrepareParams(params map[string]string) string {
	var str string = "?"
	for k, v := range params {
		str = fmt.Sprintf(str+"%s%s%s%s", k, "=", v, "&")
	}
	return str[0:strings.LastIndex(str, "&")]
}
