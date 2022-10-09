package helpers

import (
	"net/http"
)

func Get(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
