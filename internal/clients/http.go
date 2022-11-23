package clients

import "net/http"

func newClinet() *http.Client {
	return &http.Client{}
}

func Excute(request *http.Request) (*http.Response, error) {
	response, err := newClinet().Do(request)
	if err != nil {
		return &http.Response{}, err
	}
	return response, nil
}
