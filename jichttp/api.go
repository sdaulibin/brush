package jichttp

import (
	"binginx.com/brush/config"
	"binginx.com/brush/model"
	"net/http"
)

type JicHttp struct {
	client *http.Request
}

func New(user *model.UserInfo, url string, method string, params *config.Params) (*http.Request, error) {
	defaultHttpRequest := config.NewDefaultHttpRequest(user.Token, url, method, params)
	req, err := http.NewRequest(defaultHttpRequest.Method, defaultHttpRequest.Url, nil)
	req.Header = *defaultHttpRequest.Header
	return req, err
}
