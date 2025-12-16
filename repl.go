package main

import "strings"

func cleanInput(text string) []string {

	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	parts := strings.Split(lowered, " ")

	return parts
}