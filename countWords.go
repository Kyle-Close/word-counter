package main

import (
	"regexp"
	"strings"
)

func CountWords(content string) map[string]int {
	words := strings.Fields(content)
	results := make(map[string]int)

	for _, v := range words {
		key := strings.ToLower(v)
		key = removePunctuation(key)

		if strings.TrimSpace(key) != "" {
			results[key]++
		}
	}

	return results
}

func removePunctuation(str string) string {
	regex := regexp.MustCompile(`[\?!,\.]`)
	return regex.ReplaceAllString(str, "")
}
