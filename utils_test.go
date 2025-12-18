package main

import (
	"testing"
	"fmt"
)

func TestHttpJSONGet(t *testing.T) {
	cases := []struct {
		input string
		expected1 []map[string]interface{}
		expected2 error
	}{
		{
			input: "https://api.boot.dev/v1/courses_rest_api/learn-http/projects/0194fdc2-fa2f-4cc0-81d3-ff12045b73c8",
			expected1: []map[string]any{
				map[string]any{
					"assignees":6,
					"completed":false,
					"id":"0194fdc2-fa2f-4cc0-81d3-ff12045b73c8",
					"title":"Marketing Campaign Q4",
				},
			},
			expected2: nil,
		},
		{
			input: "https://api.boot.dev/v1/courses_rest_api/learn-http/projects/1234",
			expected1: nil,
			expected2: fmt.Errorf("received non-200 response code: 404"),
		},
	}

	for _, c := range cases {
		actual1, actual2 := httpJSONGet(c.input)
		if fmt.Sprint(actual1) != fmt.Sprint(c.expected1) || fmt.Sprint(actual2) != fmt.Sprint(c.expected2) {
			t.Errorf("TestHttpJSONGet failed: for input '%s' expected '%v','%v' but got '%v','%v'", c.input, c.expected1, c.expected2, actual1, actual2)
		}
	}
}