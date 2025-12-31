package main

import (
	"net/http"
	"fmt"
	"io"
)

func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}