package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func httpJSONGet(url string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", res.StatusCode)
	}

	var data any
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	switch v := data.(type) {
	case []any:
		result = make([]map[string]any, len(v))
		for i, item := range v {
			if m, ok := item.(map[string]any); ok {
				result[i] = m
			} else {
				return nil, fmt.Errorf("unexpected item type in arrayse")
			}
		}
	case map[string]any:
		return []map[string]any{v}, nil
	default:
		return nil, fmt.Errorf("unexpected item type in array")
	}

	return result, nil
}