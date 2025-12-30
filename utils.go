package main

import (
	"net/http"
	"fmt"
	"io"
)

func httpJSONGet(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", res.StatusCode)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}